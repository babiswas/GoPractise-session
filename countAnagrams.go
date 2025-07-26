package app

import "fmt"

func CountAnagrams(strin string, pat string) int {
	start := 0
	end := 0
	source := []rune(strin)
	window := len(pat)
	mp := make(map[rune]int)
	for _, val := range pat {
		cnt, ok := mp[val]
		if ok {
			mp[val] = cnt + 1
		} else {
			mp[val] = 1
		}
	}
	count := len(mp)
	counter := 0
	fmt.Println(mp)
	for end < len(source) {
		cnt, ok := mp[source[end]]
		if ok {
			mp[source[end]] = cnt - 1
			if mp[source[end]] == 0 {
				count = count - 1
			}
		}
		if (end - start + 1) == window {
			if count == 0 {
				fmt.Println(mp)
				counter += 1
			}
			cnt, ok := mp[source[start]]
			if ok {
				mp[source[start]] = cnt + 1
				if cnt == 0 {
					count += 1
				}
			}
			start += 1
		}
		end += 1
	}
	return counter
}
