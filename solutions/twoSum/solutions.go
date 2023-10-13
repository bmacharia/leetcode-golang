package twosum

// the function signature is a slice of integers and an integer that represents the target
func twoSum(nums []int, target int) []int {
	// create a map to store the distinct values and their index
	mapset := make(map[int]int)
	// iterate through the slice
	for i, v := range nums {
		// this calcutes the difference between the target and the current value
		targetValue := target - v
		//check if the value is stored in the map
		if idx, ok := mapset[targetValue]; ok {
			// if it is return the index of the value and the current index
			return []int{idx, i}
		}
		// if the value is not in the map, insert the value into the
		mapset[v] = i
	}
	// if no values are found return an empty slice
	return []int{}
}
