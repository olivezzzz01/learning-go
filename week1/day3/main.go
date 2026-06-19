package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	var name string = "Bob"
	age := 30

	score := 95
	average := float64(score)

	var active bool

	msg := fmt.Sprintf(
		"%s is %d years old. Score: %.1f",
		name,
		age,
		average,
	)

	fmt.Println(msg)
	fmt.Println("Friday =", Friday)
	fmt.Println("Zero bool =", active)

	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

}
