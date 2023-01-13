package main

import (
	"fmt"
	"untitled/problems"
)

func main() {
	fmt.Println("Hello, World!")
	numIslands()
	//lonelyInteger() //TODO:
	medianInArray()

	//AddTwoNumbers
	l1 := problems.ListNode{2, &problems.ListNode{4, &problems.ListNode{3, nil}}}
	l2 := problems.ListNode{5, &problems.ListNode{6, &problems.ListNode{4, nil}}}
	problems.AddTwoNumbers(&l1, &l2)
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
	//grid := [][]byte{
	//	{1, 1, 0, 0, 0},
	//	{1, 1, 0, 0, 0},
	//	{0, 0, 1, 0, 0},
	//	{0, 0, 0, 1, 1},
	//}

	//Case 3
	grid := [][]byte{
		{1, 1, 1, 0, 1},
		{0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{1, 0, 0, 1, 0},
	}

	//Case 4
	//grid := [][]byte{
	//	{1, 0, 1, 0, 1, 0},
	//	{0, 1, 0, 1, 0, 1},
	//	{1, 0, 1, 0, 1, 0},
	//	{0, 1, 0, 1, 0, 1},
	//	{1, 0, 1, 0, 1, 0},
	//}

	//In the Go language, you are allowed to assign a function to a variable. When
	//you assign a function to a variable, then the type of the variable is of
	//function type, and you can call that variable like a function call
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if grid[r][c] != 1 { // graph search starts with 1 in pos
			return false // return to root or recursive calls in the process of graph search
		}
		grid[r][c] = 2 //required to block already searched nodes when searching within graph traversal or from root
		//graph search below
		if r-1 >= 0 { //see that you are not in water (out of bounds)
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
		return true // return to root
	}
	n := 0
	for r := range grid {
		for c := range grid[r] {
			if dfs(r, c) { //root
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

//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//Another test case: [3,2,3]: expected [0,2]
func twoSum(nums []int, target int) []int {
	// Space: O(n)
	aNum := make(map[int]int) //much better: make(map[int]int,len(nums))
	// Time: O(n)
	for idx, num := range nums {
		// Time: O(1)
		if pos, ok := aNum[target-num]; ok {
			return []int{pos, idx}
		} else {
			aNum[num] = idx
		}
	}
	return []int{}
}
