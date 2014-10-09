package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 5, 10)
	mySlice = append(mySlice, 1, 2, 3)

	for i := 0; i < len(mySlice); i++ {
		fmt.Println("mySlice[", i, "] =", mySlice[i])
	}

	fmt.Println("len(mySlice)=", len(mySlice))
	fmt.Println("cap(mySlice)=", cap(mySlice))
}
