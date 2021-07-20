package main

import "fmt"

/*
	- Ngày: 20.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	var (
		number1 = 5.25
		number2 = 1.5
	)
	avg := average(number1, number2)
	fmt.Printf("Trung bình của số %0.2f và %0.2f là %0.3f", number1, number2, avg)
}

func average(firstNumber float64, secondNumber float64) float64 {
	return (firstNumber + secondNumber) / 2
}
