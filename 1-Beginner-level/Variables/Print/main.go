package main

import (
	"fmt"
)

func main() {

	fmt.Printf("Hello \n")
	person := "Masum"
	stringInter := fmt.Sprintf("Hello %s \n", person)

	fmt.Printf(stringInter)

	age := 26
	fmt.Printf("VariableType: Person %T Age %T\n", person, age)
	fmt.Printf("Person: %s Age: %d \n", person, age)
}
