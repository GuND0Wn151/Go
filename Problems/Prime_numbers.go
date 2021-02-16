package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the ending range : ")
	input.Scan()
	a, _ := strconv.ParseInt(input.Text(), 10, 16)
	fmt.Println(a)
	for i := int64(2); i <= a; i++ {
		k := 0
		//fmt.Println(i)
		for j := int64(2); j < i; j++ {
			if i%j == 0 {
				k++
				break
			}
		}
		if k == 0 {
			fmt.Println(i)
		}
	}
}
