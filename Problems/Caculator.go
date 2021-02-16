package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var a1, b1 int

	for true {
		fmt.Println("__________Calculator__________")
		fmt.Print("1. Addition\n2. Subtraction\n3. Multiplication\n4. division\n5. Exit\n")
		fmt.Print("Enter Your Option: ")
		input.Scan()
		a, _ := strconv.ParseInt(input.Text(), 10, 64)
		//fmt.Println(a)
		if a != 5 {
			fmt.Print("Enter The Values: ")
			fmt.Scanf("%d %d", &a1, &b1)
			switch {
			case a == 1:
				fmt.Println("Addition =", a1+b1)
			case a == 2:
				fmt.Println("Subtraction =", a1-b1)
			case a == 3:
				fmt.Println("multiplication =", a1*b1)
			case a == 4:
				if b1 != 0 {
					fmt.Println("division =", a1/b1)
				} else {
					fmt.Println("Error Division by Zero")
				}
			}
		} else {
			break
		}
		input.Scan()
	}
}
