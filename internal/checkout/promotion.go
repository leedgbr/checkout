package checkout

type buyNXGetSomeXFreePromotion struct {
	sku                string
	quantityRequired   int
	quantityDiscounted int
	unitPrice          int
}

func newBuyNXGetSomeXFreePromotion(sku string, quantityRequired int, quantityDiscounted int, unitPrice int) *buyNXGetSomeXFreePromotion {
	return &buyNXGetSomeXFreePromotion{
		sku:                sku,
		quantityRequired:   quantityRequired,
		quantityDiscounted: quantityDiscounted,
		unitPrice:          unitPrice,
	}
}

func (p buyNXGetSomeXFreePromotion) adjustment(sku []string) int {
	itemsOrdered := countItemsOrdered(sku, p.sku)
	howManyDiscountsApply := itemsOrdered / p.quantityRequired
	return -howManyDiscountsApply * (p.quantityDiscounted * p.unitPrice)
}

type buyNXGetAllXDiscountedPromotion struct {
	sku                string
	quantityRequired   int
	discountPercentage int
	unitPrice          int
}

func newBuyNXGetAllXDiscountedPromotion(sku string, quantityRequired int, discountPercentage int, unitPrice int) *buyNXGetAllXDiscountedPromotion {
	return &buyNXGetAllXDiscountedPromotion{
		sku:                sku,
		quantityRequired:   quantityRequired,
		discountPercentage: discountPercentage,
		unitPrice:          unitPrice,
	}
}

func (p buyNXGetAllXDiscountedPromotion) adjustment(sku []string) int {
	itemsOrdered := countItemsOrdered(sku, p.sku)
	if itemsOrdered < p.quantityRequired {
		return 0
	}
	// handle financial calculation in a way that will not loose accuracy with the numbers we expect for our use case.
	totalPrice := itemsOrdered * p.unitPrice
	discount := p.discountPercentage / 100
	discountRemainder := p.discountPercentage % 100
	adj := -totalPrice * discount
	adjRemainder := -totalPrice * discountRemainder / 100
	return adj + adjRemainder
}

type buyXGetYFreePromotion struct {
	sku              string
	freeSKU          string
	freeSKUUnitPrice int
}

func newBuyXGetYFreePromotion(sku string, freeSKU string, freeSKUUnitPrice int) *buyXGetYFreePromotion {
	return &buyXGetYFreePromotion{
		sku:              sku,
		freeSKU:          freeSKU,
		freeSKUUnitPrice: freeSKUUnitPrice,
	}
}

func (p buyXGetYFreePromotion) adjustment(sku []string) int {
	itemsOrdered := countItemsOrdered(sku, p.sku)
	discountedItemsOrdered := countItemsOrdered(sku, p.freeSKU)
	discountValue := 0
	for i := 0; i < itemsOrdered; i++ {
		if discountedItemsOrdered >= i+1 {
			discountValue = discountValue + p.freeSKUUnitPrice
		}
	}
	return -discountValue
}

func countItemsOrdered(sku []string, skuToFind string) int {
	itemsOrdered := 0
	for _, itemSKU := range sku {
		if skuToFind == itemSKU {
			itemsOrdered++
		}
	}
	return itemsOrdered
}

func samplePromotions(inventory map[string]*Item) []promotion {
	googleQuantityRequired := 3
	googleQuantityDiscounted := 1
	alexaQuantityRequired := 3
	alexaDiscountPercentage := 10
	return []promotion{
		newBuyNXGetSomeXFreePromotion(googleHomeSKU, googleQuantityRequired, googleQuantityDiscounted, inventory[googleHomeSKU].priceInCents),
		newBuyNXGetAllXDiscountedPromotion(alexaSKU, alexaQuantityRequired, alexaDiscountPercentage, inventory[alexaSKU].priceInCents),
		newBuyXGetYFreePromotion(macbookProSKU, raspberryPiSKU, inventory[raspberryPiSKU].priceInCents),
	}
}
