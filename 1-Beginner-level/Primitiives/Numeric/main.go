package main

import "fmt"

func main() {

	// Represent a singed integer, Which can range from -32768 to 32767
	var var1 int16
	// Represent a unsinged integer, Which can range from 0 to 65365
	var var2 uint16

	fmt.Printf("%v, %T \n", var1, var1)
	fmt.Printf("%v, %T \n", var2, var2)

	a := 10
	b := 3

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)

	// We can do any operation on same type of varible, we can't do any operation on two different type varible.
	fmt.Printf("%v, %T \n", a+b, a+b)
	fmt.Printf("%v, %T \n", a-b, a-b)
	fmt.Printf("%v, %T \n", a/b, a/b) // Here we get 3, as we are dividing a int to int so we will get a result in int.
	fmt.Printf("%v, %T \n", a*b, a*b)
	fmt.Printf("%v, %T \n", a%b, a%b) // we can only perfom this with int.

	var c float64
	c = float64(a / b)
	fmt.Printf("%v, %T \n", c, c)

	ba := 10 // binary representation - 1010 ( 1 * 2^3 + 0 * 2^2 + 1 * 2^1 + 0 * 2^0 = 10)
	bb := 3  // binary representation - 0011 ( 0 * 2^3 + 0 * 2^2 + 1 * 2^1 + 1 * 2^0 = 3)

	// Byte vise operations ( We can bit shiift for only int type.)

	fmt.Printf("%v \n", ba&bb)  // 1010 & 0011 = 0010 -> 2
	fmt.Printf("%v \n", ba|bb)  // 1010 | 0011 = 1011 --> 11
	fmt.Printf("%v \n", ba^bb)  // 1010 ^ 0011 = 1001 --> 9 // ^ Exclusive Or means either one has bit set, or the other does but not both.
	fmt.Printf("%v \n", ba&^bb) // 1010 &^ 0011 = 1000 --> 8 // &^ AND NOT means,  It clears specific bits in a value. Essentially, for each bit in the left-hand operand, if the corresponding bit in the right-hand operand is 1, the result bit is set to 0; otherwise, the result bit is the same as the left-hand operand's bit.
	// In simple word,
	// If the bit in bb is 1, the corresponding bit in ba is cleared (set to 0).
	// If the bit in bb is 0, the corresponding bit in ba remains unchanged.

	// Bit shifiting ( We can bit shiift for only int type.)

	sa := 10

	fmt.Printf("%v \n", sa<<3) // multiply bit represent with 2 ^ shift-number
	// 2^3 + 2^1 -- 2^3 * 2^3 + 2^1 * 2^3 = 80
	fmt.Printf("%v \n", sa>>3) // divide bit represent with 2 ^ shift-number
	// 2^3 + 2^1 -- 2^3 / 2^3 + 2^1 / 2^3 = 1 + 1 / 4 = 1 + 0.25 = 1.25 = 1
}
