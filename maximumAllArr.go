package app

import "container/list"

func MaxmAllSubarr(arr []int, window int) []int {
	start := 0
	end := 0
	maxlist := list.New()
	result := make([]int, 0)
	for end < len(arr) {
		if maxlist.Len() == 0 {
			maxlist.PushBack(arr[end])
		} else {
			for arr[end] > maxlist.Back().Value.(int) {
				maxlist.Remove(maxlist.Back())
				if maxlist.Len() == 0 {
					break
				}
			}
			maxlist.PushBack(arr[end])
		}
		if end-start+1 == window {
			maxvalue := maxlist.Front().Value.(int)
			result = append(result, maxvalue)
			if arr[start] == maxvalue {
				elm := maxlist.Front()
				maxlist.Remove(elm)
			}
			start += 1
		}
		end += 1
	}
	return result
}
