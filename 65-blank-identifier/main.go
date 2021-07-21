package main

import (
	"fmt"
	"math"
)

/*
	- Ngày: 21.07.2021
	- By: Nguyễn Đăng Đức
*/
func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

func main() {
	prevStockPrice := 80000.0
	currentStockPrice := 120000.0

	change, _ := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("Giá giảm khoảng $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("Giá giảm khoảng $%.2f\n", math.Abs(change))
	}
}
