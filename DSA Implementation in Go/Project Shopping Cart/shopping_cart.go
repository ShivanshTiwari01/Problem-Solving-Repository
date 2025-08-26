package main

import "fmt"

type Item struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type ShoppingCart struct {
	Items []Item
	Total float64
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		Items: make([]Item, 0),
		Total: 0.0,
	}
}

func (cart *ShoppingCart) AddItem(id int, name string, price float64, quantity int) {
	item := Item{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
	cart.Items = append(cart.Items, item)
	cart.calculateTotal()
}

func (cart *ShoppingCart) RemoveItem(id int) {
	for i, item := range cart.Items {
		if item.ID == id {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			break
		}
	}
	cart.calculateTotal()
}

func (cart *ShoppingCart) UpdateQuantity(id int, newQuantity int) {
	for i := range cart.Items {
		if cart.Items[i].ID == id {
			cart.Items[i].Quantity = newQuantity
			break
		}
	}
	cart.calculateTotal()
}

func (cart *ShoppingCart) calculateTotal() {
	cart.Total = 0.0
	for _, item := range cart.Items {
		cart.Total += item.Price * float64(item.Quantity)
	}
}

func (cart *ShoppingCart) DisplayCart() {
	fmt.Println("\n--- Shopping Cart ---")
	if len(cart.Items) == 0 {
		fmt.Println("Your cart is empty")
		return
	}
	
	for _, item := range cart.Items {
		fmt.Printf("ID: %d | %s | $%.2f x %d = $%.2f\n", 
			item.ID, item.Name, item.Price, item.Quantity, 
			item.Price*float64(item.Quantity))
	}
	fmt.Printf("\nTotal: $%.2f\n", cart.Total)
}

func main() {
	cart := NewShoppingCart()
	
	// Demo usage
	cart.AddItem(1, "Apple", 1.50, 3)
	cart.AddItem(2, "Banana", 0.75, 5)
	cart.AddItem(3, "Orange", 2.00, 2)
	
	cart.DisplayCart()
	
	fmt.Println("\nUpdating banana quantity to 8...")
	cart.UpdateQuantity(2, 8)
	cart.DisplayCart()
	
	fmt.Println("\nRemoving orange...")
	cart.RemoveItem(3)
	cart.DisplayCart()
}
