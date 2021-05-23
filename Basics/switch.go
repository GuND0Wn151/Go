package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func main(){
    input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a value: ")
	input.Scan()
	
	x,_:=strconv.ParseInt(input.Text(),10,64)
	x=x%2
	switch x{
	case 0:
		fmt.Printf("It is Even")
	case 1:
		fmt.Printf("It is ODD")
	}
}
