/*
Maximum Subarray
Easy

16812

786

Add to List

Share
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
I'm deferring this for a later time. Hopefully when I get to the algorithms section I should be able to do this.
*/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	globalMax := -100000
	localMax := -100000
	for _, v := range nums {
		localMax = max(v, v+localMax)
		if localMax > globalMax {
			globalMax = localMax
		}
	}
	return globalMax
}

// Kadane's algorithm

// Had to look into the solution to get this algorithm