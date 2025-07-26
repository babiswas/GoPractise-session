package app

func Permutation(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)

	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		Permutation(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
