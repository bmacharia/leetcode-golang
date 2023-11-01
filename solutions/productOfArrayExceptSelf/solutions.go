package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	//create a result slice that is the same length as the input slice
	res := make([]int, len(nums))
	// create a prefix variable and set it to 1, it will be used to store the product of all the numbers to the left of the current number
	prefix := 1
	// iterate through the input slice
	for i, num := range nums {
		// set the current value in the result slice to the prefix
		res[i] = prefix
		// multiply the prefix by the current number
		prefix *= num
	}
	// create a postfix variable and set it to 1, it will be used to store the product of all the numbers to the right of the current number
	postfix := 1
	// iterate through the input slice backwards,from right to left
	for i := len(nums) - 1; i >= 0; i-- {
		// multiply the current value in the result slice by the postfix
		res[i] *= postfix
		// multiply the postfix by the current number in the nums slice
		postfix *= nums[i]
	}
	return res
}
