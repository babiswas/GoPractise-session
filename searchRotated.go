package app

func SearchRotatedArr(arr []int, key int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > arr[start] {
			if (key < arr[mid]) && (key >= arr[start]) {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if (key > arr[mid]) && (key <= arr[end]) {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
