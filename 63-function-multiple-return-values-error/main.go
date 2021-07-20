package main

import (
	"errors"
	"fmt"
	"math"
)

/*
	- Ngày: 20.07.2021
	- By: Nguyễn Đăng Đức
*/
func getStockPriceChangeWithError(oldPrice, newPrice float64) (float64, float64, error) {
	if oldPrice == 0 {
		err := errors.New("Giá cũ không được để = 0")
		return 0, 0, err
	}
	change := newPrice - oldPrice
	percentChange := (change / oldPrice) * 100
	return change, percentChange, nil
}

func main() {
	oldPrice := 0.0
	newPrice := 100000.0

	change, percentChange, err := getStockPriceChangeWithError(oldPrice, newPrice)

	if err != nil {
		fmt.Println("Có lỗi xảy ra => ", err)
	} else {
		if change < 0 {
			fmt.Printf("Giá mới giảm $%.2f với tỉ lệ thay đổi %.2f%% so với giá cũ\n", math.Abs(change), math.Abs(percentChange))
		} else {
			fmt.Printf("Giá mới tăng $%.2f với tỉ lệ thay đổi %.2f%% so với giá cũ\n", math.Abs(change), math.Abs(percentChange))

		}
	}
}
