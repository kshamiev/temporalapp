package activities

import (
	"context"

	"temporalapp/internal/generated/temporal"
)

func (a *Activities) AssortmentReserveCancel(ctx context.Context, req *temporal.AssortmentReserveRequest) error {
	return a.assortmentClient.ReserveCancel(ctx, req.Products)
}
