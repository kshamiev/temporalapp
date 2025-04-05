package checkout

import (
	"go.temporal.io/sdk/workflow"

	"temporalapp/internal/generated/temporal"
	"temporalapp/internal/utils"
)

func Register(ctx workflow.Context, input *temporal.CheckoutFlowWorkflowInput) (temporal.CheckoutFlowWorkflow, error) {
	return &Workflow{
		req: input.Req,
	}, nil
}

type Workflow struct {
	req *temporal.CheckoutFlowRequest
}

func (w *Workflow) Execute(ctx workflow.Context) (*temporal.Order, error) {
	// Резервируем продукты
	reserveProducts := make([]*temporal.AssortmentReserveProductRequest, 0, len(w.req.Cart.Products))
	for _, p := range w.req.Cart.Products {
		reserveProducts = append(reserveProducts, &temporal.AssortmentReserveProductRequest{
			Id:  p.Id,
			Qty: p.Qty,
		})
	}
	err := temporal.AssortmentReserve(ctx, &temporal.AssortmentReserveRequest{Products: reserveProducts})
	if err != nil {
		return nil, err
	}

	// Пример компенсации, здесь больше и сложнее:
	// https://github.com/temporalio/samples-go/blob/main/saga/workflow.go#L35
	defer func() {
		// Отменяем резерв
		if err != nil {
			// Если не получится отменить резерв, то лучше куда-то эскалировать или иметь фолбэк
			_ = temporal.AssortmentReserveCancel(ctx, &temporal.AssortmentReserveRequest{
				Products: reserveProducts,
			})
		}
	}()

	if w.req.PaymentType == temporal.PaymentType_ONLINE {
		p, err := temporal.CreatePayment(ctx, &temporal.CreatePaymentRequest{})
		if err != nil {
			return nil, err
		}
		// Пример компенсации, здесь больше и сложнее:
		// https://github.com/temporalio/samples-go/blob/main/saga/workflow.go#L35
		defer func() {
			// Отменяем платеж
			if err != nil {
				// Если не получится отменить резерв, то лучше куда-то эскалировать или иметь фолбэк
				_ = temporal.PaymentCancel(ctx, &temporal.PaymentCancelRequest{
					Id: p.Id,
				})
			}
		}()
	}
	order := &temporal.Order{
		Id:          utils.WorkflowID(ctx),
		Customer:    w.req.Customer,
		Cart:        w.req.Cart,
		PaymentType: w.req.PaymentType,
	}

	// Здесь мы запускаем дочернее Workflow, но уже с политикой Abandon
	// https://docs.temporal.io/encyclopedia/child-workflows#parent-close-policy
	//workflowcheck:ignore
	run, err := temporal.ProcessingFlowChildAsync(ctx, &temporal.ProcessingFlowRequest{
		Id:          order.Id,
		Customer:    order.Customer,
		Cart:        order.Cart,
		PaymentType: order.PaymentType,
	})
	if err != nil {
		return nil, err
	}
	_, err = run.WaitStart(ctx)
	if err != nil {
		return nil, err
	}

	return order, nil
}
