package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Welcome to Yasith's first Go program. Hi, what is your name? ")
	fmt.Scanln(&name)

	fmt.Print("How old are you? ")
	fmt.Scanln(&age)

	fmt.Println("Hi, your name is", name, "and you are", age, "years old.\nThank you for using this program. Goodbye...")
}
