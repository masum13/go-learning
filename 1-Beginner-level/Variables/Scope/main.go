package main

import "fmt"

// Global/ Package level variable declaration which can we avilable to all
var (
	person  = "Masum"
	age     = 26
	gender  = "Male"
	married = false
)

func main() {
	fmt.Printf("Person details\nName: %s Age: %d Gender: %s Married: %t \n", person, age, gender, married)

	PrintMsg() // TODO: See difrenece on each function call.

	// We can change the variable value by define inside the function level,
	age = 27
	fmt.Printf("Person updated details\nName: %s Age: %d Gender: %s Married: %t \n", person, age, gender, married)

	// We can declare more function level variable

	email := "masum@gmail.com"
	fmt.Printf("person updated details\nName: %s Age: %d Gender: %s Married: %t Email: %s\n", person, age, gender, married, email)

	married = true
	if married {
		var Massage = "Congrets !!"
		fmt.Printf("%s %s \n", Massage, person) // Output: I'm a variable inside a block
	}

	// We can't be able to access Message variable outside of block, We will get undefined: MassagecompilerUndeclaredName.
	// fmt.Printf("%s %s \n", Massage, person)

	PrintMsg()
}

func PrintMsg() {
	// We can't be able to access email variable here as it is define at a function level. Error: undefined: EmailcompilerUndeclaredName
	// fmt.Printf("Person updated details\nName: %s Age: %d Gender: %s Married: %t Email: %s\n", Person, age, gender, Married, email)
	customMsg := "Verifed Account."
	fmt.Printf("Person details\nName: %s Age: %d Gender: %s Married: %t Massage: %s \n", person, age, gender, married, customMsg)
}
