package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
