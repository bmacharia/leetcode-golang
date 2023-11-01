package groupanagrams

func Groupanagrams(strs []string) [][]string {
	// create a map of [26]int to []string
	// the [26]int represents the count of each letter in the string
	// it will be the key to group the anagrams together
	anagramMap := make(map[[26]int][]string)
	// itereate through the string values in the input slice
	for _, s := range strs {
		// initilize a fixed size arrary of 26 ints
		// at this point we are creating a key for this string value that will be used to insert into the map
		var count [26]int
		//interate through the string value character by character
		for _, c := range s {
			// for each character c increment the corresponding index in the count array
			count[c-'a']++
		}
		// append the string s to the slice in the map that has a matching key
		anagramMap[count] = append(anagramMap[count], s)
	}
	// initilize a slice of slices of strings
	var result [][]string
	// for each group of anagrams in the map append the group to the result slice
	for _, group := range anagramMap {
		// append the group to the result slice
		result = append(result, group)
	}
	// return the result slice of slices that represents the groups of anagrams
	return result
}
