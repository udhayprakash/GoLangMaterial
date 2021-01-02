package main

import "fmt"

func main() {

	stockPrices := map[string]float32{
		"IBM":  85.23,
		"HPQ":  399.00,
		"AAPL": 712.00,
		"MSFT": 123.00,
	}

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Original map content")
	fmt.Println("---------------------------------------------------------------------------")

	for k, v := range stockPrices {
		fmt.Printf("key : %v, value : %v \n", k, v)
	}

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("New map subset with prices higher than 200")
	fmt.Println("---------------------------------------------------------------------------")

	higherStockPrices := map[string]float32{}

	// copy map value with prices higher than 200
	// to new map

	for name, price := range stockPrices {
		if price > 200.00 {
			fmt.Println(name, price)
			higherStockPrices[name] = price
		}
	}

	//fmt.Println(higherStockPrices)

	for higherName, higherPrice := range higherStockPrices {
		fmt.Printf("key : %v , value : %v \n", higherName, higherPrice)
	}

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("New map subset with stock names that are found in slice")
	fmt.Println("---------------------------------------------------------------------------")

	stockNames := []string{"HPQ", "MSFT"}

	for name, price := range stockPrices {
		for _, sName := range stockNames {
			if name == sName {
				fmt.Println(name, price)
			}
		}
	}

}
