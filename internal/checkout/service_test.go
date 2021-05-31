package checkout

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testAlexaSKU       = "A304SD"
	testGoogleHomeSKU  = "120P90"
	testMacbookProSKU  = "43N23P"
	testRaspberryPiSKU = "234234"
)

func TestCheckout(t *testing.T) {
	type test struct {
		name          string
		sku           []string
		expectedTotal int
	}

	tests := []test{
		{"empty cart", []string{}, 0},
		{"google home", []string{testGoogleHomeSKU}, 4999},
		{"macbook pro", []string{testMacbookProSKU}, 539999},
		{"alexa speaker", []string{testAlexaSKU}, 10950},
		{"raspberry pi", []string{testRaspberryPiSKU}, 3000},
		{"google home promotion", []string{testGoogleHomeSKU, testGoogleHomeSKU, testGoogleHomeSKU}, 9998},
		{"alexa speaker promotion", []string{testAlexaSKU, testAlexaSKU, testAlexaSKU, testAlexaSKU}, 39420},
		{"raspberry pi promotion", []string{testMacbookProSKU, testRaspberryPiSKU}, 539999},
	}

	s := newSimpleService()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, _ := s.checkout(tc.sku)
			assert.Equal(t, &result{TotalInCents: tc.expectedTotal}, r)
		})
	}
}

func TestCheckoutProductNotFound(t *testing.T) {
	s := newSimpleService()
	r, err := s.checkout([]string{"*bad-sku*"})
	assert.Nil(t, r)
	assert.Equal(t, fmt.Errorf("item not found"), err)
}
