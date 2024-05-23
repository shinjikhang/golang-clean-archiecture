package simpleinterest

import "fmt"

// Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years

func init() {
	fmt.Println("simpleinterest init 01")
}

func init() {
	fmt.Println("simpleinterest init 02")
}

func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
