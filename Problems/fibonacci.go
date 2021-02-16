package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number of elements: ")
	input.Scan()
	a, _ := strconv.ParseInt(input.Text(), 10, 64)
	f := 0
	k := 0
	b := 1
	for i := int64(0); i < a; i++ {
		fmt.Println(f)
		k = f
		f = b
		b = b + k
	}

}
