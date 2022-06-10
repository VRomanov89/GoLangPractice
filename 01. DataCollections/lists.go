package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[0])

	featuredPrices := prices[1:3] // Slice of an array
	fmt.Println(featuredPrices)

}
