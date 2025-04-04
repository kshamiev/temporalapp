package activities

import (
	"temporalapp/internal/services/assortment"
	"temporalapp/internal/services/payment"
)

type Activities struct {
	assortmentClient *assortment.Client
	paymentClient    *payment.Client
}

func Register(ac *assortment.Client, pc *payment.Client) *Activities {
	return &Activities{
		assortmentClient: ac,
		paymentClient:    pc,
	}
}
