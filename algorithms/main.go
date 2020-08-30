package main

import (
	"github.com/slashpai/computer_science/algorithms/searching"
	"fmt"
)

func main() {
	list :=[]int {2, 3, 4, 5, 8, 9, 20, 25, 39 }
	result := searching.Search(list, 0, len(list) - 1, 1)
	fmt.Println("Element found? :", result)
}
