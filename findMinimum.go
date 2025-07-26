package app

func MinimumElement(arr []int) int {
	start := 0
	end := len(arr) - 1
	for start < end {
		if arr[start] < arr[end] {
			return start
		}
		mid := (start + end) / 2
		prev := (mid - 1 + len(arr)) % len(arr)
		end := (mid + 1) % len(arr)
		if arr[mid] < arr[prev] && arr[mid] < arr[end] {
			return mid
		} else if arr[mid] >= arr[start] {
			start = mid + 1
		} else if arr[mid] >= arr[end] {
			end = mid - 1
		}
	}
	return -1
}
