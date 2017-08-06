package supermarket

import "fmt"

func main() {
	fmt.Println("--- Supermarket Price ---")
	var basket = make([]ProductSelection, 0)
	basket = append(basket, ProductSelection{Amt: 3, Prod: Product{Price: 10, Name: "Meat"}})
	basket = append(basket, ProductSelection{Amt: 4, Prod: Product{Price: 20, Name: "Cheese"}})

	b := Basket{contents: basket}

	fmt.Println("One Free Discount=", CalculatePrice(b, OneFreeDiscountStrategy))
	fmt.Println("No Discount=", CalculatePrice(b, NoDiscountStrategy))
}
