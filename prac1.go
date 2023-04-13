package main

import (
	"fmt"
	"math"
)

const s string = "Hello,there!"

func main() {
	var Pname = "Ayush Rusia"
	var Lname = "Yoooo"
	var Sname = "port"
	fmt.Printf("Pname: %s", Pname)
	fmt.Printf("\nLname: %s", Lname)
	fmt.Printf("\nSname: %s\n", Sname)

	var X int16 = 24
	var Y int16 = 48
	fmt.Printf("addition : %d + %d = %d\n", X, Y, X+Y)
	fmt.Printf("subtraction : %d - %d = %d\n", Y, X, Y-X)
	fmt.Printf("multiplication : %d * %d = %d\n", Y, X, Y*X)
	fmt.Printf("division : %d / %d = %d\n", Y, X, Y/X)

	a := 30.23
	b := 45.24
	c := b - a
	fmt.Printf("Subtration of %f and %f is %f\n", b, a, c)
	fmt.Printf("The type of %f is %T\n", c, c)

	var str1 = "AyushRusia"
	var str2 = "ayushrusia"
	var str3 = "AyushRusia"
	result1 := str1 == str2
	result2 := str1 == str3

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("The type of results is %T\n", result1)

	var str4 = "HelloWorld"
	fmt.Printf("The type of string 4 is %T\n", str4)
	fmt.Printf("The length of string 4 is %d\n", len(str4))

	var str5 = "Ayush_"
	var str6 = "Rusia"
	fmt.Printf("The new string is %s \n", str5+str6)

	var num int
	var str7 string
	var fl float64
	fmt.Printf("The value of num is %d\n", num)
	fmt.Printf("The value of str7 is %s\n", str7)
	fmt.Printf("The value of floating point number is %f\n", fl)

	myvar1, myvar2 := 24, 50
	myvar3, myvar2 := 23, 100
	fmt.Printf("The value of myvar1 is %d\n", myvar1)
	fmt.Printf("The value of myvar2 is %d\n", myvar2)
	fmt.Printf("The value of myvar3 is %d\n", myvar3)

	const PI = 3.14
	fmt.Printf("The value of PI is %f\n", PI)

	const trueConst = true

	type myBool bool
	var defaultBool bool = trueConst
	var customBool myBool = trueConst

	fmt.Printf("The value of defaultBool is %t\n", defaultBool)
	fmt.Printf("The value of customBool is %t\n", customBool)

	fmt.Println(s)
	const n = 5
	const d = 3e10 / n
	fmt.Println(d)
	fmt.Printf("The value of d is %d\n", int64(d))
	fmt.Printf("The value of n is %f\n", math.Sin(n))

	var v int = 1200
	if v < 1000 {
		fmt.Printf("v is less than 1000\n")
	} else {
		fmt.Printf("v is greater than 1000\n")
	}

	for i := 0; i < 4; i++ {
		fmt.Println("AyushRusia")
	}

}
