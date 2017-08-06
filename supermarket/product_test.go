package supermarket

import (
	"fmt"
	"testing"
)

func TestNoDiscount(t *testing.T) {
	fmt.Println("--- Supermarket Price Test ---")
	var basket = make([]ProductSelection, 0)
	basket = append(basket, ProductSelection{Amt: 3,
		Prod: Product{Price: 10, Name: "Meat"}})
	basket = append(basket, ProductSelection{Amt: 4,
		Prod: Product{Price: 20, Name: "Cheese"}})

	b := Basket{contents: basket}

	response := CalculatePrice(b, NoDiscountStrategy)
	if response != 110 {
		t.Error("Expected 110")
	}
}
