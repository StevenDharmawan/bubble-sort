package main

import "fmt"

func main() {
	intArray := []int{7, 2, 1, 4, 9, 10, 8, 5, 6, 3}
	//fmt.Println(intArray)
	sortedArray := bubbleSort2(intArray)
	fmt.Println(sortedArray)
}

func bubbleSort2(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			//fmt.Printf("i = %d j = %d\n", i, j)
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				fmt.Println(array)
			}
			//else {
			//	fmt.Println("SKIP")
			//}
		}
	}
	return array
}
