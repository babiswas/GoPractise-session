package app

func LastOccurance(brr []int, element int) int {
	start := 0
	end := len(brr) - 1
	result := 0
	for start <= end {
		mid := (start + end) / 2
		if brr[mid] == element {
			result = mid
			start = mid + 1
		} else if brr[mid] > element {
			end = mid - 1
		} else if brr[mid] < element {
			start = mid + 1
		}
	}
	return result
}
