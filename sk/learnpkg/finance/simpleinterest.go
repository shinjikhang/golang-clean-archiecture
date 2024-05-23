package finance

import (
	"fmt"
)

func init() {
	fmt.Println("finance init 01")
}

func init() {
	fmt.Println("finance init 02")
}

func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
