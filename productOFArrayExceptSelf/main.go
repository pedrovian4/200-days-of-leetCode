package productOFArrayExceptSelf

func _(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	for i := range answer {
		answer[i] = 1
	}

	leftProduct := 1
	rightProduct := 1

	for i := 0; i < length; i++ {
		answer[i] *= leftProduct
		leftProduct *= nums[i]
	}

	for i := length - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer
}
