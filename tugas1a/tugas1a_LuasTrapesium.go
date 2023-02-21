package main

import "fmt"

func main() {
	var a, b, t, luas float64

	fmt.Print("Masukkan nilai a : ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b : ")
	fmt.Scan(&b)

	fmt.Print("Masukkan nilai t : ")
	fmt.Scan(&t)

	luas = 0.5 * (a + b) * t
	fmt.Println("Luas Trapesium anda adalah ", luas)

}
