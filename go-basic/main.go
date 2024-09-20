package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0)
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Append: ", i)
		fmt.Printf("Address: %p\n", &mySlice)
		fmt.Printf("Len: %d\n", len(mySlice))
		fmt.Printf("Cap: %d\n", cap(mySlice))
		fmt.Println("================================")
	}
}
