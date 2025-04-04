package activities

import (
	"context"

	"temporalapp/internal/generated/temporal"
)

func (a *Activities) AssortmentReserve(ctx context.Context, req *temporal.AssortmentReserveRequest) error {
	return a.assortmentClient.Reserve(ctx, req.Products)
}
