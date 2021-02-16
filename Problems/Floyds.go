package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number of lines: ")
	input.Scan()
	k := int64(1)
	a, _ := strconv.ParseInt(input.Text(), 10, 64)
	for i := int64(0); i < a; i++ {
		for j := int64(0); j <= i; j++ {
			fmt.Print(k, " ")
			k++
		}
		fmt.Println()
	}
}
