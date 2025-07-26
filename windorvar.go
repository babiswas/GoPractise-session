package app

func FindWindowSumK(brr []int, K int) []int {
	start := 0
	end := 0
	result := make([]int, 0)
	sum := 0
	for end < len(brr) {
		sum += brr[end]
		if sum == K {
			result = append(result, end-start+1)
		} else if sum > K {
			temp := start
			tempSum := sum
			for temp != end {
				tempSum -= brr[temp]
				if tempSum == K {
					result = append(result, end-temp)
				}
				temp += 1
			}
		}
		end += 1
	}
	return result
}
