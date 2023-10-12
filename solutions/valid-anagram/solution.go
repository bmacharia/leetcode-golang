package validanagram

func isAnagram(s string, t string) bool {
	// check if the length of the input strings are equal
	if len(s) != len(t) {
		return false
	}
	// create a counter array to keep
	// track of the occurence of each character
	var counter []int

	// iterate over the input strings
	// increment the counter array based on index for each character in s
	// decrement the counter array based on index for each character in t
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}
	// iterate over the counter array
	// if any of the values are not 0, return false
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	// return true if all the values in the counter array are 0
	return true
}
