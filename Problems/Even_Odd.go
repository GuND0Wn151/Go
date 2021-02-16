package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number: ")
	input.Scan()
	a, _ := strconv.ParseInt(input.Text(), 10, 64)
	if a%2 == 0 {
		fmt.Println(a, "is Even")
	} else {
		fmt.Println(a, "is Odd")
	}

}
