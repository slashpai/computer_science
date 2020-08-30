package searching;

import(
	"fmt"
)

func Search(list []int, start, end, key int) bool {
	fmt.Println(list, start, end, key)
  if start < 0 || end >= len(list) || start > end {
		fmt.Println("Boundary Condition!")
		return false
	}
	mid := (start + end) / 2
	fmt.Println("mid index = ", mid)
	fmt.Println("mid value = ", list[mid])
	if list[mid] == key {
		return true
	} else if key < list[mid] {
			end = mid - 1
  	} else if key > list[mid] {
			start = mid + 1
		}
		return Search(list, start, end, key)
	}
