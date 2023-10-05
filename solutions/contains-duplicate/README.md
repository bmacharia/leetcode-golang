# LeetCode Problem 217. Contains Duplicate

[Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)

## Thought Process to Solve the Problem

One way to solve this problem is to use a hash set to keep track of the unique elements in the array. We can iterate through the array and for each element, we can check if it already exists in the hash set. If it does, we can return true as we have found a duplicate element. If we reach the end of the array without finding a duplicate element, we can return false


You can call this function with an integer array as an argument to check if it contains any duplicate elements.

## Time Complexity

The time complexity of the hash set solution to the LeetCode problem 217 "Contains Duplicate" is O(n), where n is the length of the input array. This is because we iterate through the array once and perform constant time operations (insertion and lookup) on the hash set for each element. Therefore, the time taken to solve the problem is proportional to the size of the input array.