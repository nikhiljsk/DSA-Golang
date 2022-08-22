// DO NOT SUBMIT !!
// This is a naive approach and timesout if submitted as it is of order n-square
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
*/

// DO NOT SUBMIT !!
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		tempSum2 := nums[i]
		tempSum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			tempSum += nums[j]
			if tempSum > tempSum2 {
				tempSum2 = tempSum
			}
		}
		if tempSum2 > maxSum {
			maxSum = tempSum2
		}
	}
	return maxSum
}

// Able to crack the approach and implemented
// DO NOT SUBMIT !!