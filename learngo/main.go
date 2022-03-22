package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood1 := []string{"chicken", "pizza", "hamburger"}
	kim := person{"kim", 20, favFood1}
	fmt.Println(kim)

	favFood2 := []string{"kimchi", "namul"}
	choi := person{name: "choi", favFood: favFood2}
	fmt.Println(choi)
}
