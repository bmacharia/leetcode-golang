# LeetCode Problem 242. Valid Anagram

[Valid Anagram](https://leetcode.com/problems/valid-anagram/)

## Thought Process to Solve the Problem

To solve this problem to check if the two input strings are anagrams of each other. An anagram is a word or phrase formed by rearranging the letters of another word or phrase. In this case, we are checking if the two input strings have the same set of characters, regardless of their order.

The function takes two string arguments, `s` and `t`, and returns a boolean value. The first thing the function does is check if the lengths of the two input strings are equal. If they are not equal, the function immediately returns `false`, since two strings of different lengths cannot be anagrams.

If the lengths of the two input strings are equal, the function creates an integer slice called `counter` to keep track of the frequency of each character in the two input strings. The slice has a length of 26, which corresponds to the number of lowercase letters in the English alphabet.

The function then iterates over each character in the `s` string and increments the corresponding value in the `counter` slice. It does the same thing for the `t` string, but decrements the corresponding value in the `counter` slice.

Finally, the function iterates over the values in the `counter` slice and checks if any of them are non-zero. If any of the values are non-zero, the function returns `false`, since the two input strings are not anagrams. If all of the values are zero, the function returns `true`, since the two input strings are anagrams.



## Time Complexity

Overall, this function has a time complexity of O(n), where n is the length of the input strings. This is because we iterate over each character in each string once and perform constant time operations (increment and decrement) on the `counter` slice for each character. Therefore, the time taken to solve the problem is proportional to the size of the input strings.



