package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert an integer to float64
	var num int = 10
	var numFloat float64 = float64(num)
	numFloat01 := float64(num)
	fmt.Printf("Converted float: %f\n", numFloat)
	fmt.Printf("Converted float: %f\n", numFloat01)

	// Convert a float64 to int
	var floatNum float64 = 3.14
	var intNum int = int(floatNum) // Here we will lose the information, So Use with caution.
	fmt.Printf("Converted integer: %d\n", intNum)

	// Convert a string to []byte (byte slice)
	var str string = "Hello, GoLang!"
	var byteSlice []byte = []byte(str)
	fmt.Printf("Converted byte slice: %v\n", byteSlice)

	// Convert int to string, We can't do that directly but we can use strconv package to do that.
	var numStr int
	numStr = 40
	strNum := strconv.Itoa(numStr) // This will convert int to string.
	fmt.Printf("Type: numStr %T strNum %T \n", numStr, strNum)

	// We can also do the reverse with strconv.Atoi function. But use with caution .
	numStrNew, _ := strconv.Atoi(strNum)
	fmt.Printf("Type: strNum %T numStrNew %T \n", strNum, numStrNew)

}
