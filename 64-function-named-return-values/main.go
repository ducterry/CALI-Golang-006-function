package main

import (
	"fmt"
	"math"
)

/*
	- Ngày: 20.07.2021
	- By: Nguyễn Đăng Đức
*/
func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return change, percentChange
}

func main() {
	prevStockPrice := 100000.0
	currentStockPrice := 90000.0

	change, percentChange := getNamedStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("Giá mới giảm $%.2f với tỉ lệ thay đổi %.2f%% so với giá cũ\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("Giá mới tăng $%.2f với tỉ lệ thay đổi %.2f%% so với giá cũ\n", math.Abs(change), math.Abs(percentChange))

	}
}
