package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	medianInArray()
}

//medianInArray give the median of a given array. Constraints length is always
//odd, -10000 <= arr[i] <= 10000
func medianInArray() {
	arr := []int32{2, 1, 5, 3, 4}
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	fmt.Println(arr)
	fmt.Println(arr[(l-1)/2])
}
