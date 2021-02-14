package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	a, _ := strconv.ParseInt(input.Text(), 10, 64)
	x := a + 1
	//this is normal while kinda for loop
	//which has a condition it repeats the for loop until the condition is true
	//we have to update the condtion is x++ or x-- or else it will run in infinite loop
	for x != 0 { //conditions
		fmt.Println(x)
		x-- //updation
	}
	// this is the for loop syntax
	//which has 3 parst initilization condition and updation
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if a == 4 {
			fmt.Println(" a = ", i)
			break
		}
	}

}
