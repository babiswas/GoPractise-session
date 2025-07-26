package app

func Permute(brr []string, start int, arr *[]string) {
	if start == len(brr) {
		temp := ""
		for _, sample := range brr {
			temp += sample
		}
		*arr = append(*arr, temp)
	}
	for i := start; i < len(brr); i++ {
		brr[start], brr[i] = brr[i], brr[start]
		Permute(brr, start+1, arr)
		brr[start], brr[i] = brr[i], brr[start]
	}
}
