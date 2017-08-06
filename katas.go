package main

import "fmt"
import "supermarket"

func main() {
	fmt.Println("--- Supermarket Price ---")
	var basket = make([]supermarket.ProductSelection, 0)
	basket = append(basket, supermarket.ProductSelection{Amt: 3, Prod: supermarket.Product{Price: 10, Name: "Meat"}})
	basket = append(basket, supermarket.ProductSelection{Amt: 4, Prod: supermarket.Product{Price: 20, Name: "Cheese"}})

	b := supermarket.Basket{Contents: basket}

	fmt.Println("One Free Discount=", supermarket.CalculatePrice(b, supermarket.OneFreeDiscountStrategy))
	fmt.Println("No Discount=", supermarket.CalculatePrice(b, supermarket.NoDiscountStrategy))
}
