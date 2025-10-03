package usecases

import "github.com/ngoctb13/forya-be/internal/domains/supply/repos"

type Supply struct {
	supplyRepo repos.ISupply
}

func NewSupply(supplyRepo repos.ISupply) *Supply {
	return &Supply{supplyRepo: supplyRepo}
}
