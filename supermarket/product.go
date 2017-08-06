package supermarket

import "math"

// Product is
type Product struct {
	Price float64
	Name  string
}

// ProductSelection is
type ProductSelection struct {
	Amt  int
	Prod Product
}

// Basket is
type Basket struct {
	contents []ProductSelection
}

// --- Price strategies ---

func CalculatePrice(b Basket, strategy func(Basket) float64) float64 {
	return strategy(b)
}

// NoDiscountStrategy No discount at all
func NoDiscountStrategy(b Basket) float64 {
	sum := 0.0
	for _, item := range b.contents {
		sum += float64(item.Amt) * item.Prod.Price
	}
	return sum
}

// OneFreeDiscountStrategy reduce one item per type in the basket
func OneFreeDiscountStrategy(b Basket) float64 {
	var newContents = make([]ProductSelection, 0)
	for _, item := range b.contents {
		amt := int(math.Max(1, float64(item.Amt-1)))
		item.Amt = amt
		newContents = append(newContents, item)
	}
	newB := Basket{contents: newContents}
	return NoDiscountStrategy(newB)
}
