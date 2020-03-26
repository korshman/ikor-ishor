package main

import ( 
	"fmt" 
	"strconv" 
)
func main() {
	var decimal int64
	fmt.Print("When you enter Decimal Number: ")
	fmt.Scanln(&decimal)
	output := strconv.FormatInt(decimal, 2)
	fmt.Print("Its equivalent in Binary is: ", output)
}