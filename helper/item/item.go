package item

import (
	"math/rand"

	"../../helper"
)

// SetItemName used to generate item name
func SetItemName(id int) string {
	var itemName string

	switch id {
	case 0:
		itemName = "Bronze"
	case 1:
		itemName = "Silver"
	case 2:
		itemName = "Gold"
	case 3:
		itemName = "Platinum"
	}

	return itemName
}

// StorageAvailable used to check available storage
// for either player or broker
func StorageAvailable(quantity int) bool {
	return quantity >= 0 && quantity < 100
}

// SetInitialPrice is used to generate initial item prices
// IDs
// 0 -> Bronze (Price within 200-600)
// 1 -> Silver (Price within 600-1000)
// 2 -> Gold (Price within 1000-1400)
// 3 -> Platinum (Price within 1400-2000)
func SetInitialPrice(id int) int {
	var itemPrice int

	switch id {
	case 0:
		itemPrice = helper.GenerateRandomInteger(200, 600)
	case 1:
		itemPrice = helper.GenerateRandomInteger(600, 1000)
	case 2:
		itemPrice = helper.GenerateRandomInteger(1000, 1400)
	case 3:
		itemPrice = helper.GenerateRandomInteger(1400, 2000)
	}

	return itemPrice
}

// UpdatePrice is used to update prices based on probabiity
// 50% chances prices can change or else
// If change occur, 70% price increase and 30% else
// Prices can raise from 0-25%, or drop from 0-15% range
func UpdatePrice(price int) int {
	var priceUpdateChance int
	var priceChangeChance int

	priceUpdateChance = rand.Intn(100)
	priceChangeChance = rand.Intn(100)

	if priceUpdateChance < 50 {
		if priceChangeChance < 70 {
			price = price + price*rand.Intn(25)/100
		} else {
			price = price - price*rand.Intn(15)/100
		}
	}

	return price
}

// UpdateQuantity used for updating quantity of an item stock
func UpdateQuantity(status int, currentQuantity int, changeQuantity int) int {
	if status == 1 {
		currentQuantity = currentQuantity + changeQuantity
	} else {
		currentQuantity = currentQuantity - changeQuantity
	}

	return currentQuantity
}
