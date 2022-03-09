package main

import "fmt"

func main() {
	var action int

	fmt.Println("Enter 1 for Student or 2 for Professional.")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println("I am a student.")
	case 2:
		fmt.Println("I am a professional.")
	default:
		panic(fmt.Sprintf("Invalid input %d", action))
	}

	defer func() {
		action := recover()
		fmt.Println("BOO!")
		fmt.Println(fmt.Sprintf("Recover returned %s", action))
	}()
}
