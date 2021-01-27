package main

import "fmt"

func main() {
	//normal declaration of variables
	var var1 int = 10

	//declaration of variables without specifying data type
	//in this type go complier automatically chosses the appropriate datatype
	var var2 = "New String"

	fmt.Println(var1, var2)

	//There is also another way in which we can declare useing the assingment operator
	// ":=" is known as assingment operator and also kown as walrus operator
	// In this no need of using keyword "var" and also no need of mentioning data type

	//String variable
	var3 := "New Varible string"

	// int variable
	var4 := 100

	//printing values
	fmt.Println(var3, var4)

	// In go to know the datatype of the variable we can use %T
	fmt.Printf("%T", var2)
	fmt.Printf(" %T", var1)

	// here printf is similar to formating string %t refers to type of variable

}
