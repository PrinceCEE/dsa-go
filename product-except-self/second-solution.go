package main

func productExceptSelf(nums []int) []int {
	arrLen := len(nums)

	product := make([]int, arrLen)
	suffix := make([]int, arrLen)
	prefix := make([]int, arrLen)

	prefix[0] = 1
	for i := 1; i < arrLen; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix[arrLen-1] = 1
	for i := arrLen - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < arrLen; i++ {
		product[i] = prefix[i] * suffix[i]
	}

	return product
}
