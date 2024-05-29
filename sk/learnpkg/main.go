package main

import (
	"fmt"
	//_ "learnpackage/finance"
	//_ "learnpackage/simpleinterest"
)

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	salary := employeeSalary["joe"]
	fmt.Println("Salary of jamie", "is", salary)
}

func init() {
	fmt.Println("----------Main initialized first----------")
}
