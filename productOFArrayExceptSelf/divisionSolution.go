package productOFArrayExceptSelf

func _(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	totalProduct := 1
	zeroCount := 0
	for _, num := range nums {
		if num != 0 {
			totalProduct *= num
		} else {
			zeroCount++
		}
	}

	for i, num := range nums {
		if zeroCount == 0 {
			answer[i] = totalProduct / num
		} else if zeroCount == 1 && num == 0 {
			answer[i] = totalProduct
		} else {
			answer[i] = 0
		}
	}

	return answer
}
