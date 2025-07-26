package app

func FirstOccurance(arr []int, element int) int {
	start := 0
	last := len(arr) - 1
	temp := 0
	for start <= last {
		mid := (start + last) / 2
		if arr[mid] == element {
			temp = mid
			last = mid - 1
		} else if arr[mid] > element {
			last = mid - 1
		} else if arr[mid] < element {
			start = mid + 1
		}
	}
	return temp
}
