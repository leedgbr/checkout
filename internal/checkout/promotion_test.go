package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	discountedSKU = "discounted-sku"
	freeSKU       = "free-sku"
	sku           = "sku"
)

func TestBuyXGetSomeXFree(t *testing.T) {
	type test struct {
		name               string
		sku                []string
		expectedAdjustment int
	}

	tests := []test{
		{"0 items", []string{}, 0},
		{"0 discounted items", []string{sku}, 0},
		{"1 discounted item", []string{discountedSKU}, 0},
		{"2 discounted items", []string{discountedSKU, discountedSKU}, 0},
		{"3 discounted items", []string{discountedSKU, discountedSKU, discountedSKU}, -500},
		{"4 discounted items", []string{discountedSKU, discountedSKU, discountedSKU, discountedSKU}, -500},
		{"5 discounted items", []string{discountedSKU, discountedSKU, discountedSKU, discountedSKU, discountedSKU}, -500},
		{"6 discounted items", []string{discountedSKU, discountedSKU, discountedSKU, discountedSKU, discountedSKU, discountedSKU}, -1000},
		{"7 discounted items", []string{discountedSKU, discountedSKU, discountedSKU, discountedSKU, discountedSKU, discountedSKU, discountedSKU}, -1000},
		{"7 discounted items with other products", []string{sku, discountedSKU, discountedSKU, discountedSKU, discountedSKU, sku, discountedSKU, discountedSKU, discountedSKU, sku}, -1000},
	}

	p := newBuyNXGetSomeXFreePromotion(discountedSKU, 3, 1, 500)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAdjustment, p.adjustment(tc.sku))
		})
	}
}

func TestBuyXGetYFree(t *testing.T) {
	type test struct {
		name               string
		sku                []string
		expectedAdjustment int
	}

	tests := []test{
		{"0 items", []string{}, 0},
		{"0 discounted 0 free items", []string{sku}, 0},
		{"1 discounted 0 free items", []string{discountedSKU}, 0},
		{"1 discounted 1 free item", []string{discountedSKU, freeSKU}, -500},
		{"1 discounted 2 free item", []string{discountedSKU, freeSKU, freeSKU}, -500},
		{"1 discounted 3 free item", []string{discountedSKU, freeSKU, freeSKU, freeSKU}, -500},
		{"2 discounted 3 free item", []string{discountedSKU, discountedSKU, freeSKU, freeSKU, freeSKU}, -1000},
		{"3 discounted 3 free item", []string{discountedSKU, discountedSKU, discountedSKU, freeSKU, freeSKU, freeSKU}, -1500},
		{"4 discounted 3 free item", []string{discountedSKU, discountedSKU, discountedSKU, discountedSKU, freeSKU, freeSKU, freeSKU}, -1500},
		{"4 discounted 3 free item with other products", []string{sku, discountedSKU, discountedSKU, sku, discountedSKU, discountedSKU, freeSKU, freeSKU, freeSKU, sku}, -1500},
	}

	p := newBuyXGetYFreePromotion(discountedSKU, freeSKU, 500)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAdjustment, p.adjustment(tc.sku))
		})
	}
}

func TestBuyNXGetAllXDiscounted(t *testing.T) {
	type test struct {
		name               string
		sku                []string
		expectedAdjustment int
	}

	tests := []test{
		{"0 items", []string{}, 0},
		{"0 discounted", []string{sku}, 0},
		{"1 discounted", []string{discountedSKU}, 0},
		{"2 discounted", []string{discountedSKU, discountedSKU}, 0},
		{"3 discounted", []string{discountedSKU, discountedSKU, discountedSKU}, -300},
		{"4 discounted", []string{discountedSKU, discountedSKU, discountedSKU, discountedSKU}, -400},
		{"5 discounted", []string{discountedSKU, discountedSKU, discountedSKU, discountedSKU, discountedSKU}, -500},
		{"5 discounted with other products", []string{sku, discountedSKU, sku, discountedSKU, discountedSKU, discountedSKU, discountedSKU, sku}, -500},
	}

	p := newBuyNXGetAllXDiscountedPromotion(discountedSKU, 3, 20, 500)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAdjustment, p.adjustment(tc.sku))
		})
	}

}

func TestBuyNXGetAllXDiscountedFractionalCents(t *testing.T) {
	sku := []string{discountedSKU, discountedSKU}
	p := newBuyNXGetAllXDiscountedPromotion(discountedSKU, 1, 1, 100000000)
	assert.Equal(t, -2000000, p.adjustment(sku))
}
