package customer

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"temporalapp/internal/actions"
	"temporalapp/internal/generated/temporal"
	"temporalapp/internal/utils"
)

func Register(ctx workflow.Context, input *temporal.CustomerFlowWorkflowInput) (temporal.CustomerFlowWorkflow, error) {
	return &Workflow{
		profile: &temporal.Profile{
			Id:    utils.WorkflowID(ctx),
			Name:  input.Req.Name,
			Phone: input.Req.Phone,
		},
		deleteProfileSignal: input.DeleteProfile,
		deleteCartSignal:    input.DeleteCart,
		setAddressSignal:    input.SetAddress,
	}, nil
}

type Workflow struct {
	profile             *temporal.Profile
	cart                *temporal.Cart
	deleteProfileSignal *temporal.DeleteProfileSignal
	deleteCartSignal    *temporal.DeleteCartSignal
	setAddressSignal    *temporal.SetAddressSignal
}

func (w *Workflow) Execute(ctx workflow.Context) error {
	// Ожидаем 1 минуту:
	//	- пока пользователь введет адрес
	//  - или отменяем его "жизненный цикл"
	var isCancelled bool
	sel := workflow.NewSelector(ctx)
	w.setAddressSignal.Select(sel, func(request *temporal.SetAddressRequest) {
		w.profile.Address = request.Address
	})
	sel.AddFuture(workflow.NewTimer(ctx, time.Minute), func(f workflow.Future) {
		isCancelled = true
	})
	sel.Select(ctx)
	if isCancelled {
		return workflow.ErrCanceled
	}

	// Создаем логер
	logger := workflow.GetLogger(ctx)

	// Задаем стандартные настройки для activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: time.Second})

	// Выполняем HelloActivity
	var helloResult string
	if err := workflow.ExecuteActivity(ctx, actions.HelloActivity, "HelloActivity").Get(ctx, &helloResult); err != nil {
		logger.Error("HelloActivity failed.", "Error", err)
		return err
	}
	logger.Info("HelloActivity activity completed.", "result", helloResult)

	// Выполняем ByeActivity
	var byeResult string
	if err := workflow.ExecuteActivity(ctx, actions.ByeActivity, "ByeActivity").Get(ctx, &byeResult); err != nil {
		logger.Error("ByeActivity failed.", "Error", err)
		return err
	}
	logger.Info("ByeActivity activity completed.", "result", byeResult)

	// Дальше запускаем бесконечный цикл:
	// - обрабатываем обновление адреса
	// - удаление профиля (с завершением цикла)
	// - удаление корзины
	// событие обновления корзины и создания заказа живут в отдельных хэндлерах
	var stop bool
	for !stop {
		sel := workflow.NewSelector(ctx)
		w.setAddressSignal.Select(sel, func(request *temporal.SetAddressRequest) {
			w.profile.Address = request.Address
		})
		w.deleteProfileSignal.Select(sel, func() {
			stop = true
		})
		w.deleteCartSignal.Select(sel, func() {
			w.cart = nil
		})
		sel.Select(ctx)
	}

	return nil
}

func (w *Workflow) GetProfile() (*temporal.Profile, error) {
	return w.profile, nil
}

func (w *Workflow) GetCart() (*temporal.Cart, error) {
	return w.cart, nil
}

func (w *Workflow) UpdateProfile(ctx workflow.Context, request *temporal.UpdateProfileRequest) (*temporal.Profile, error) {
	w.profile.Name = request.Name
	return w.profile, nil
}
