# LeetCode Problem 49 Group Anagrams

[GroupAnagram](https://leetcode.com/problems/group-anagrams/)

## Thought Process to Solve the Problem

The function works by creating a map where the key is an array of 26 integers that represent the count of each letter in a string. This key is used to group anagrams together. The function iterates through each string in the input slice and creates a key for it by initializing a fixed-size array of 26 integers and incrementing the corresponding index in the array for each character in the string. The string is then appended to the slice in the map that has a matching key.

After all the strings have been processed, the function initializes a slice of slices of strings and iterates through each group of anagrams in the map. For each group, the function appends the group to the result slice. Finally, the function returns the result slice of slices that represents the groups of anagrams.





## Time Complexity

The time complexity of the selected code is O(n * k), where n is the length of the input slice and k is the maximum length of a string in the input slice. 

This is because the function iterates through each string in the input slice and creates a key for it by initializing a fixed-size array of 26 integers and incrementing the corresponding index in the array for each character in the string. This takes O(k) time for each string. 

After all the strings have been processed, the function iterates through each group of anagrams in the map and appends the group to the result slice. This takes O(n) time since there can be at most n groups of anagrams. 

Therefore, the overall time complexity of the function is O(n * k).
