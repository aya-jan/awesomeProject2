package main

import (
	"github.com/aya-jan/awesomeProject2/models"
	"time"
)

func main() {
	shoppingCart := models.NewShoppingCart(getItems())
	shoppingCart.ShowItems()
	shoppingCart.Add(getItem())
	shoppingCart.ShowItems()
	shoppingCart.Remove(0)
	shoppingCart.ShowItems()

	visa := models.NewCreditCardStrategy("Ayazhan", "5522043350849269", "123", time.Now())
	shoppingCart.Pay(&visa)
	shoppingCart.ShowItems()

	shoppingCart = models.NewShoppingCart(getItems())
	shoppingCart.ShowItems()
	shoppingCart.Add(getItem())
	shoppingCart.ShowItems()
	shoppingCart.Remove(0)
	shoppingCart.ShowItems()
	paypal := models.NewPayPalStrategy("ayazhan@gmail.com", "password!234")
	shoppingCart.Pay(&paypal)
	shoppingCart.ShowItems()
}

func getItems() []models.Item {
	var items []models.Item
	items = append(items, models.Item{
		Price: 101.0,
		Code:  "BREAD",
	})

	items = append(items, models.Item{
		Price: 50.0,
		Code:  "EGG",
	})

	items = append(items, models.Item{
		Price: 999.99,
		Code:  "CHICKEN",
	})

	items = append(items, models.Item{
		Price: 240.0,
		Code:  "TOMATO",
	})
	return items
}

func getItem() models.Item {
	return models.Item{
		Price: 12391.11,
		Code:  "RANDOM_ITEM",
	}
}
