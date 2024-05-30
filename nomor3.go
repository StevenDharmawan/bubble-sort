package main

import "fmt"

func main() {
	//intArray := []int{7, 2, 1, 4, 9, 10, 8, 5, 6, 3}
	intArray := []int{1, 2, 3, 4, 5}
	//fmt.Println(intArray)
	sortedArray := bubbleSort3(intArray)
	fmt.Println(sortedArray)
}

func bubbleSort3(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		counter := 0
		for j := 0; j < len(array)-i-1; j++ {
			//fmt.Printf("i = %d j = %d\n", i, j)
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				fmt.Println(array)
				counter++
			}
			//else {
			//	fmt.Println("SKIP")
			//}
		}
		if counter == 0 {
			break
		}
	}
	return array
}
