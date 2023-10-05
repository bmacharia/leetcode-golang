// this function detects whether a slice of integers contains any duplicates
// if it does the fucntion returns true, otherwise it returns false (no duplicates)

package containsduplicate

func ContainsDuplicate(nums []int) bool {
	// create a map to store the numbers that stores distinct numbers
	set := make(map[int]bool)
	// iterate over the nums slice
	for _, num := range nums {
		// if the number is already in the map, return true
		// this is a lookup o
		if set[num] {
			return true
		}
		// otherwise, add the number to the map
		// this is an insert operation to add the element to the map
		set[num] = true
	}
	// if we get here, there are no duplicates
	return false
}
