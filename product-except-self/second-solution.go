package main

func productExceptSelf(nums []int) []int {
	arrLen := len(nums)

	product := make([]int, arrLen)
	suffix := make([]int, arrLen)
	prefix := make([]int, arrLen)

	for i := 0; i < arrLen; i++ {
		if i == 0 {
			prefix[i] = 1
		} else {
			prefix[i] = prefix[i-1] * nums[i-1]
		}
	}
	for i := arrLen - 1; i >= 0; i-- {
		if i == arrLen-1 {
			suffix[arrLen-i-1] = 1
		} else {
			suffix[arrLen-i-1] = suffix[arrLen-i-2] * nums[i+1]
		}
	}

	for i := 0; i < arrLen; i++ {
		product[i] = prefix[i] * suffix[arrLen-1-i]
	}

	return product
}
