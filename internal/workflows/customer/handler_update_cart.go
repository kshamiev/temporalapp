package customer

import (
	"errors"
	"fmt"
	"path"
	"runtime"

	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"

	"temporalapp/internal/generated/temporal"
)

func (w *Workflow) UpdateCart(ctx workflow.Context, request *temporal.UpdateCartRequest) (*temporal.Cart, error) {
	// Формируем список продуктов, которые надо получить из ассортимента
	productIDs := make([]string, len(request.Products))
	for i := range request.Products {
		productIDs[i] = request.Products[i].GetId()
	}

	// Получаем те самые продукты c помощью activity
	assortmentGetProductsReq := &temporal.AssortmentGetProductsRequest{
		Ids: productIDs,
	}
	assortment, err := temporal.AssortmentGetProducts(ctx, assortmentGetProductsReq)
	if err != nil {
		return nil, err
	}

	// Проверяем, что все продукты есть и их хватает для добавления в корзину
	stocks := make(map[string]*temporal.AssortmentProduct)
	for _, p := range assortment.Products {
		stocks[p.Id] = p
	}
	products := make([]*temporal.Product, 0, len(request.Products))
	for _, p := range request.Products {
		assortmentProduct, ok := stocks[p.Id]
		if !ok {
			return nil, errors.New("the product is not in stock")
		}
		if assortmentProduct.Stocks < p.Qty {
			return nil, errors.New("not enough goods in stock")
		}
		products = append(products, &temporal.Product{
			Id:    assortmentProduct.Id,
			Name:  assortmentProduct.Name,
			Price: assortmentProduct.Price,
			Inn:   assortmentProduct.Inn,
			Qty:   p.Qty,
		})
	}

	w.cart = &temporal.Cart{
		Products: products,
		Total:    calculateTotal(products),
	}

	// Изменение с использованием версий
	switch version(ctx, 1) {
	case workflow.DefaultVersion: // старая версия кода
	case 1: // новая версия кода
		encodedValue := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return uuid.NewString()
		})
		if err := encodedValue.Get(&w.cart.Id); err != nil {
			return nil, err
		}
	}

	//workflowcheck:ignore
	{
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1")
		fmt.Println("workflow ID:", workflow.GetInfo(ctx).WorkflowExecution.ID)
		fmt.Println("workflow instance ID:", workflow.GetInfo(ctx).WorkflowExecution.RunID)
		fmt.Println("workflow instance method ID:", workflow.GetCurrentUpdateInfo(ctx).ID)
		fmt.Println()
	}

	return w.cart, nil
}

func calculateTotal(products []*temporal.Product) int32 {
	var total int32
	for i := range products {
		total += products[i].Qty * products[i].Price
	}
	return total
}

func version(ctx workflow.Context, newVersion int) workflow.Version {
	method := "method"
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			method = path.Base(fn.Name())
		}
	}
	return workflow.GetVersion(ctx, method+"-"+workflow.GetCurrentUpdateInfo(ctx).ID, workflow.DefaultVersion, workflow.Version(newVersion))
}
