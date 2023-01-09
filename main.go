package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	numIslands()
	//lonelyInteger() //TODO:
	//medianInArray()
}

func numIslands() {
	//Case 1
	//grid := [][]byte{
	//	{1, 1, 1, 1, 0},
	//	{1, 1, 0, 1, 0},
	//	{1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0},
	//}

	//Case 2
	grid := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	//In the Go language, you are allowed to assign a function to a variable. When
	//you assign a function to a variable, then the type of the variable is of
	//function type, and you can call that variable like a function call
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if grid[r][c] != 1 {
			return false
		}
		grid[r][c] = 2
		if r-1 >= 0 {
			dfs(r-1, c)
		}
		if r+1 < len(grid) {
			dfs(r+1, c)
		}
		if c-1 >= 0 {
			dfs(r, c-1)
		}
		if c+1 < len(grid[r]) {
			dfs(r, c+1)
		}
		return true
	}
	n := 0
	for r := range grid {
		for c := range grid[r] {
			if dfs(r, c) {
				n++
			}
		}
	}
	fmt.Println(n)
}

//a=[1, 1, 2] 2 is lonely. length is always odd, and it is guaranteed there is
//always one lonely integer.
func lonelyInteger() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}
	l := len(a)
	for i := 0; i < l; i++ {
		if i < 0 || i >= l {
			continue
		}
		for j := i + 1; j < l; j++ {
			if j < 0 || j >= l {
				continue
			}
			if a[j] == a[i] {

			}
		}
	}
	fmt.Println(a)
}

//medianInArray will give the median of a given array. Constraints length is always
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
