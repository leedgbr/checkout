package checkout

type Item struct {
	name     string
	priceInCents    int
	quantity int
}

var (
	googleHomeSKU  = "120P90"
	macbookProSKU  = "43N23P"
	alexaSKU       = "A304SD"
	raspberryPiSKU = "234234"
)

func sampleInventory() map[string]*Item {
	return map[string]*Item{
		googleHomeSKU:  {name: "Google Home", priceInCents: 4999, quantity: 10},
		macbookProSKU:  {name: "MacBook Pro", priceInCents: 539999, quantity: 5},
		alexaSKU:       {name: "Alexa Speaker", priceInCents: 10950, quantity: 10},
		raspberryPiSKU: {name: "Raspberry Pi B", priceInCents: 3000, quantity: 2},
	}
}
