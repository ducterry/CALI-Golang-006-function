package main

import (
	"fmt"
	"math"
)

/*
	- Ngày: 20.07.2021
	- By: Nguyễn Đăng Đức
*/
func getStockPriceChange(oldPrice, newPrice float64) (float64, float64) {
	change := newPrice - oldPrice
	percentChange := (change / oldPrice) * 100
	return change, percentChange
}

func main() {
	oldPrice := 10000.0
	newPrice := 100000.0

	change, percentChange := getStockPriceChange(oldPrice, newPrice)

	if change < 0 {
		fmt.Printf("Giá mới giảm $%.2f với tỉ lệ thay đổi %.2f%% so với giá cũ\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("Giá mới tăng $%.2f với tỉ lệ thay đổi %.2f%% so với giá cũ\n", math.Abs(change), math.Abs(percentChange))

	}
}
