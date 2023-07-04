package main

func productExceptSelf(nums []int) []int {
	var answer []int
	for i := range nums {
		product := 1
		for j, k := range nums {
			if i != j {
				product *= k
			}
		}
		answer = append(answer, product)
	}

	return answer
}
