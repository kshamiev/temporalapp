package activities

import (
	"temporalapp/internal/services/assortment"
)

type Activities struct {
	assortmentClient *assortment.Client
}

func Register(ac *assortment.Client) *Activities {
	return &Activities{
		assortmentClient: ac,
	}
}
