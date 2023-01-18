package problems

func DpCaller() {
	nums := []int{7, 2, 5, 10, 8}
	k := 2
	splitArray(nums, k)
}
func splitArray(nums []int, m int) int {
	maxElement, sumArr := 0, 0
	for _, n := range nums {
		sumArr += n
		if n > maxElement {
			maxElement = n
		}
	}

	l, r := maxElement, sumArr
	res := r
	for l <= r {
		mid := l + (r-l)/2
		if canSplit(nums, m, mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func canSplit(nums []int, m int, largest int) bool {
	curSum := 0
	subArray := 0
	for _, n := range nums {
		curSum += n
		if curSum > largest {
			subArray++
			curSum = n
		}
	}
	return subArray+1 <= m
}
