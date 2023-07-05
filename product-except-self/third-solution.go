package main

func productExceptSelf(nums []int) []int {
	arrLen := len(nums)

	answer := make([]int, arrLen)
	product := 1
	for i := 0; i < arrLen; i++ {
		answer[i] = product
		product *= nums[i]
	}

	suffix := 1
	for i := arrLen - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}
