package app

import "fmt"

func FirstNegative(arr []int, result *[]int, window int) {
	start := 0
	end := 0
	temp := make([]int, 0)
	index := 0

	for end < len(arr) {
		if arr[end] < 0 {
			temp = append(temp, end)
		}
		if end-start+1 == window {
			if len(temp) > 0 {
				index = temp[0]
				*result = append(*result, arr[index])
			} else {
				*result = append(*result, 0)
			}
			start += 1
			if start > index && len(temp) > 0 {
				temp = temp[1:]
			}
		}
		end += 1
	}
}

func FirstNegativeV2(arr []int, result *[]int, window int) {
	fmt.Println("The given array is:", arr)
	start := 0
	end := 0
	temp := make([]int, 0)
	for end < len(arr) {
		if arr[end] < 0 {
			temp = append(temp, arr[end])
		}
		if (end - start + 1) == window {
			if len(temp) > 0 {
				*result = append(*result, temp[0])
			} else {
				*result = append(*result, 0)
			}
			if arr[start] < 0 && len(temp) > 0 {
				temp = temp[1:]
			}
			start += 1
		}
		end += 1
	}
}
