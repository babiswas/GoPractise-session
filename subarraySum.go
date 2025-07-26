package app

func MaxSumSubarr(brr []int, size int) int {
	sum := 0
	maxSum := 0
	for i := 0; i < len(brr)-2; i++ {
		for j := i; j < i+3; j++ {
			sum += brr[j]
		}
		if sum > maxSum {
			maxSum = sum
			sum = 0
		}
	}
	return maxSum
}

func SMaxSubarray(brr []int, window int) int {
	start := 0
	end := 0
	sum := 0
	maxSum := 0
	for end < len(brr) {
		sum += brr[end]
		if end-start+1 == window {
			if sum > maxSum {
				maxSum = sum
			}
			sum -= brr[start]
			start += 1
		}
		end += 1
	}
	return maxSum
}
