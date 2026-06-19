package main

import "fmt"

func main() {

	//for loop type 1

	for j := 1; j <= 20; j++ {
		fmt.Println(j)
	}

	//for loop type 2

	j := 21
	for j <= 40 {
		fmt.Println(j)
		j++
	}

	//for loop type 3
	for {
		fmt.Println("Hello")
		break
	}

	age := 81

	if age < 18 {
		fmt.Println("minor")
	} else if age < 18 && age >= 21 {
		fmt.Println("old but can't drink")
	} else {
		fmt.Println("adult and can drink")
	}
	//switch cases with fallthrough
	switch {
	case age < 18:
		fmt.Println("minor")

	case age < 18 && age >= 21:
		fmt.Println("old but can't drink")

	default:
		fmt.Println("adult and can drink")
	}

	//defer
	defer fmt.Println("hello")
	fmt.Println("world")
	fmt.Println("Olivia")

}
