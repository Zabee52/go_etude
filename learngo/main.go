package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("lenAndUpper() complete") // return 이후에 동작한다.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper("kim")
	fmt.Println(totalLength, uppercase)
}
