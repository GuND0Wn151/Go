package main

import (
	// buffer input and output
	"bufio"
	"fmt"

	//os module
	"os"
	// for converting string to interger parseint
	"strconv"
)

func main() {
	// this will create a scanner object just like the scanner object in java util
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter Your Name: ")
	// To take the input fromthe console
	// use this command this will store the value in the input

	input.Scan()
	x := input.Text() // use input.text() to get the input from it
	fmt.Printf("Enter Your Age: ")

	// some times if we want to integer we can use the below commands

	input.Scan()
	// now to convert the string into integer we use this command
	//parseint take 3 arguments first one is string 2'nd is base value third is size

	y, _ := strconv.ParseInt(input.Text(), 10, 64) // here _ is taken for any errors

	fmt.Println(x, "is", y, "years old") // printing the outputs
}
