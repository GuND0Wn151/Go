package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Go also has the all arthematic operators like + - * / %
	// when we are doing a mathematical operation we should see that both the operaands are of same type
	// or else Go will thorugh a error

	// when we are performing arthematic operations with a type of operends then the output will be also of same type
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the First Integer: ")
	input.Scan()
	x1, _ := strconv.ParseInt(input.Text(), 10, 64)

	fmt.Printf("Enter the Second Integer: ")
	input.Scan()
	x2, _ := strconv.ParseInt(input.Text(), 10, 64)

	// here both x1 and x2 are integer so we can perform operation
	// for float values we can take strconv.ParseFloat(input.Text(),10,64)

	fmt.Println("Addition", x1+x2)
	fmt.Println("Substraction", x1-x2)
	// here in division when 2 numbers are integer then output will be also integer
	fmt.Println("Division", x1/x2) // here b cannot be 0 or else we will get a runtime error
	fmt.Println("Modulas", x1%x2)  // here also b cannot be 0 we get a divide by zero eroor
	fmt.Println("multiplication", x1*x2)

}
