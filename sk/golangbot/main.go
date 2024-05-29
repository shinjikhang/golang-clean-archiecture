package main

import "fmt"

func main() {
	var price = 100
	var no = 5
	var totalPrice = calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
}

func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}
