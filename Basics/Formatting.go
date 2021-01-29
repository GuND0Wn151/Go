package main

import "fmt"

func main() {

	// %S or %s for printing string
	// \n for new line \t for tabspaces
	// %d for integer %g for floats
	fmt.Printf("%s\n", "String")

	// %v for printing values
	fmt.Printf("%v\n", 10)

	// %T for printing the type of the value
	fmt.Printf("%T\n", "This is a String")

	// For printing other system values
	// %b base2
	// %o base8
	// %d base10
	// %x base16
	fmt.Printf("%b %o %d %x systems\n", 10, 20, 30, 40)

	// Padding and rounding off
	// %nd where n is the no of padding
	// %.nf where n is the no of digits you want to round off
	fmt.Printf("%10d\n", "padded")
	fmt.Printf("%-10d", "leftpadding")
	fmt.Printf("%.2f\n", 2.342342343)
	fmt.Printf("%8.2f\n padded round off", 2.3423423423) //rounding off plus padding with 10 spaces

	//padding with specific values %fnd where f is value n is size
	fmt.Printf("%07d ", 2323)
	// %e for scientific notaion like 10 e**2

	//sprintf making string from formatting
	str := fmt.Sprintf("hello %.1f %s", 2.21342, "noice")
	fmt.Println(str)

}
