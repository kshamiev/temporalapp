package activities

import (
	"context"

	"temporalapp/internal/generated/temporal"
)

func (a *Activities) CreatePayment(ctx context.Context, req *temporal.CreatePaymentRequest) (*temporal.CreatePaymentResponse, error) {
	return a.paymentClient.CreatePayment(ctx, req)
}
