package activities

import (
	"context"

	"temporalapp/internal/generated/temporal"
)

func (a *Activities) CreateVendorOrder(ctx context.Context, req *temporal.CreateVendorOrderRequest) (*temporal.CreateVendorOrderResponse, error) {
	return a.vendorsClient.CreateOrder(ctx, req)
}
