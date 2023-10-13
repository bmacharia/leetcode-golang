# Leetcode 1 Two Sum
[Two Sum](https://leetcode.com/problems/two-sum/description/)

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Thought Process to Solve the Problem

Create a map called `mapset` to store the distinct values in the input array and their indices.

Iterates through the input array using a `for` loop and calculates the difference between the target value and the current value in the array. 

If the difference is already stored in the `mapset` map, the function returns a slice containing the indices of the two values. If the difference is not in the map, the function inserts the current value and its index into the map.

If no two values are found that add up to the target value, the function returns an empty slice.

## Time Complexity

Overall, this function has a time complexity of O(n), where n is the length of the input array. This is because we iterate over each element in the array once and perform constant time operations (insertion and lookup) on the map for each element. Therefore, the time taken to solve the problem is proportional to the size of the input array.

The function can be used to solve the LeetCode problem 1 "Two Sum" and other similar problems that require finding pairs of numbers that add up to a target value.





