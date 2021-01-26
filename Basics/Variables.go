package main

import "fmt"

//we use uint for unsinged integers and in unit also we have types like
//uint8 uint16 uint32 etc

//we use int for singed integers
//here also we can find types like int8 int16 etc

func main() {

	var var1 string = "This is a string"
	var var2 uint = 100
	var var3 int = 200

	fmt.Println("Var1= ", var1)

	fmt.Println("Var2 = ", var2, "Var3 =", var3)
  
	//go also has boolean datatypes
	//true and false
  
	var var4 bool = true
	fmt.Println(var4)
}
