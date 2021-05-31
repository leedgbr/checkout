package checkout

import (
	"errors"
)

type result struct {
	TotalInCents int `json:"total"`
}

type promotion interface {
	adjustment(sku []string) int
}

func newSimpleService() *simpleService {
	inv := sampleInventory()
	return &simpleService{
		inventory:  inv,
		promotions: samplePromotions(inv),
	}
}

type simpleService struct {
	inventory  map[string]*Item
	promotions []promotion
}

func (s simpleService) checkout(sku []string) (*result, error) {
	totalInCents := 0
	for _, itemSKU := range sku {
		item := s.inventory[itemSKU]
		if item == nil {
			return nil, errors.New("item not found")
		}
		totalInCents += item.priceInCents
	}
	return &result{TotalInCents: totalInCents + s.adjustment(sku)}, nil
}

func (s simpleService) adjustment(sku []string) int {
	adjustment := 0
	for _, promotion := range s.promotions {
		adjustment = adjustment + promotion.adjustment(sku)
	}
	return adjustment
}
