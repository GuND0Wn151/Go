package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Print("Enter Your Age = ")
	fmt.Scanf("%d", &x)
	fmt.Println(x)
	fmt.Println("Is X less than 10 ? ", x < 10)
	//we can also use > for greater than something

	fmt.Println("Is X == 0 ?", x == 0)
	//we can also use ! to check if 2 values are not equal

	fmt.Println("Is X less than or equal to 10", x <= 10)
	// we can also use greater than equal to >=

	y := 10
	// OR operator symbol is || return true if anyone of the value is true
	fmt.Println("OR ", x > 10 || y > 10)

	// AND operator return true if both values are true symbol is &&
	fmt.Println("And", x > 10 && y > 10)

	// if there are situations where we can excute something based upon a condition we use if else statments
	// if statements are excueted when the contion for the if statment satifies
	// syntax if condition{ statments }
	// if we want to check multiple conditions and excuete statments based upon them then we can use else if statements

	if x > 18 {
		fmt.Println("You can vote")
	} else if x > 10 {
		fmt.Println("wait for", 18-x, "years")
	} else {
		fmt.Println("You cannot Vote")

	}
	// else statement is a stement when the if statment doesnt satify
	// for a else statment there must be a if statement or a else if statement
}
