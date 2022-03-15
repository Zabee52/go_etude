package main

import (
	"fmt"
)

func main() {
	names_slice := []string{"kim", "park", "choi"}
	addr := &names_slice[0]
	fmt.Println(names_slice[0], ":", addr)

	names_slice = append(names_slice, "chang")
	addr = &names_slice[0]
	fmt.Println(names_slice[0], ":", addr)
}
