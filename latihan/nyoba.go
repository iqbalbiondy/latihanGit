package main

import "fmt"

func main() {
	// Looping Over String

	var input string
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		fmt.Printf(string(input[i])+"-")
	}

	fmt.Println("   ----->")

	for pos,char := range input{
		fmt.Printf("Karakter %c berada dalam indeks ke %d\n",char,pos)
	}

	fmt.Println("Hallo")
	fmt.Println("p")

	

}