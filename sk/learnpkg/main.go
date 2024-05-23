package main

import (
	"fmt"
	"math/rand"
	//_ "learnpackage/finance"
	//_ "learnpackage/simpleinterest"
)

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	fmt.Println("Main package initialized")
}

func main() {
randloop:
	for {
		randNumber := rand.Intn(100)
		switch {
		case randNumber%2 == 0:
			fmt.Printf("Generated even number %d", randNumber)
			break randloop
		}
	}
}
