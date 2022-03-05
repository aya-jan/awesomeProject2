package models

import "fmt"

type ShoppingCart struct {
	Items []Item
}

type Item struct {
	Price float64
	Code  string
}

type makePayment interface {
	makePay(amount float64) bool
}

func NewShoppingCart(items []Item) (result ShoppingCart) {
	result.Items = items
	return
}

func (obj *ShoppingCart) Pay(payment makePayment) {
	if payment.makePay(obj.Calculate()) {
		fmt.Println("Payed. Now your cart is empty")
		obj.Items = []Item{}
	}
}

func (obj *ShoppingCart) Calculate() (ans float64) {
	for _, v := range obj.Items {
		ans += v.Price
	}
	return
}

func (obj *ShoppingCart) Add(item Item) {
	obj.Items = append(obj.Items, item)
}

func (obj *ShoppingCart) Remove(pos int) {
	obj.Items = append(obj.Items[:pos], obj.Items[pos+1:]...)
}

func (obj *ShoppingCart) ShowItems() {
	if len(obj.Items) == 0 {
		fmt.Println("Your Shopping Cart is empty!")
	}
	for _, v := range obj.Items {
		fmt.Println(v)
	}
	fmt.Println("-----------")
}
