package main

import (
	"fmt"
	"github.com/slashpai/computer_science/algorithms/searching"
)

func main() {
	list := []int{2, 3, 4, 5, 8, 9, 20, 25, 39}
	result := searching.BinarySearch(list, 0, len(list)-1, 8)
	fmt.Println("Element found? :", result)
	fmt.Println("---------------------------------------")
	result = searching.BinarySearch(list, 0, len(list)-1, 25)
	fmt.Println("Element found? :", result)
	fmt.Println("---------------------------------------")
	result = searching.BinarySearch(list, 0, len(list)-1, 1)
	fmt.Println("Element found? :", result)
	fmt.Println("---------------------------------------")
	result = searching.BinarySearch(list, 0, len(list)-1, 45)
	fmt.Println("Element found? :", result)
	fmt.Println("---------------------------------------")
}
