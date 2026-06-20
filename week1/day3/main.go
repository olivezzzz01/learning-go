package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Olivia"

	if name == "Olivia" {
		fmt.Println(name + " is a good girl")
	}

	//that was short declaration using :=, lets use var and get to know about nasic data types

	var age int
	age = 20

	fmt.Println(age)
	var b bool = true
	fmt.Println(b)
	var x float64 = 10.25
	fmt.Println(x)
	var s string = "Olivia"
	fmt.Println(s)

	//shunyaaataaaaaaaaaaa ok

	var z float64
	var k int
	fmt.Println(z)
	fmt.Println(k)
	fmt.Println("zero value of bool:", b)
	fmt.Println("zero value of string:", s)

	//go doesnt allow type conversions the way python does :3
	num1 := 21
	fmt.Println(num1)
	fmt.Println(float64(num1))

	var o int
	o = 9
	y := 2.5
	fmt.Println(float64(o) + y) //hmmm..besh shob type ek lagbe taholei hobe

	const pi = 3.14159
	var r float64
	r = 10
	fmt.Println("Area is ", pi*r*r)

	//iota concept, completely new to me :D
	// by default, iota increases by 1. so multiplying by 2 increases by 2
	// multiplying iota by any 'n' increases counter by n
	// adding any 'x' to iota increases counter by x. can set starting index
	// using this added offset.

	const (
		Monday = (iota * 2) + 10
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Saturday)
	fmt.Println(Sunday)

	//String manipulation~ upper case conv, len, concat.

	var myStr string = "I love mohiner ghoraguli"
	fmt.Println(myStr)
	fmt.Println(strings.ToUpper(myStr))

	myStr2 := "Ghore Pherar Gaan soothes"

	fmt.Println(len(myStr))
	fmt.Println(myStr + " <3 " + myStr2)

	message := fmt.Sprintf("My name is %s and I am %d years old", name, age)
	fmt.Println(message)

	//feeling complete now, can sleep peacefully :p

}
