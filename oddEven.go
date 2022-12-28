package main

import "fmt"

func main() {
	fmt.Print("Enter Your Number : ")

	var data int

	fmt.Scanln(&data)
	if data%2 == 0 {
		fmt.Println("Your Number is Even")
	} else {
		fmt.Println("Your Number is Odd")
	}
}
