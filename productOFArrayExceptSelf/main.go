package productOFArrayExceptSelf

func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	for i := range answer {
		answer[i] = 1
	}

	leftProduct := 1
	rightProduct := 1

	// Calculate the product of elements on the left side of each index
	for i := 0; i < length; i++ {
		answer[i] *= leftProduct
		leftProduct *= nums[i]
	}

	// Calculate the product of elements on the right side of each index
	for i := length - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer
}
