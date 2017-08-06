package supermarket

import (
	"fmt"
	"testing"
)

func TestNoDiscount(t *testing.T) {
	fmt.Println("--- TestNoDiscount ---")
	var basket = make([]ProductSelection, 0)
	basket = append(basket, ProductSelection{Amt: 3,
		Prod: Product{Price: 10, Name: "Meat"}})
	basket = append(basket, ProductSelection{Amt: 4,
		Prod: Product{Price: 20, Name: "Cheese"}})

	b := Basket{Contents: basket}

	response := CalculatePrice(b, NoDiscountStrategy)
	if response != 110 {
		t.Error("Expected 110, got=", response)
	}
}

func TestOneFreeDiscountStrategy(t *testing.T) {
	fmt.Println("--- TestOneFreeDiscountStrategy ---")
	var basket = make([]ProductSelection, 0)
	basket = append(basket, ProductSelection{Amt: 3,
		Prod: Product{Price: 10, Name: "Meat"}})
	basket = append(basket, ProductSelection{Amt: 4,
		Prod: Product{Price: 20, Name: "Cheese"}})

	b := Basket{Contents: basket}

	response := CalculatePrice(b, OneFreeDiscountStrategy)
	if response != 80 {
		t.Error("Expected 80, got=", response)
	}
}
