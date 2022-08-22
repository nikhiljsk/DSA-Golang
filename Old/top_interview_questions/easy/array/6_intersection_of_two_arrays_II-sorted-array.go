/*
Follow up question: 1. What if the given array is already sorted? How would you optimize your algorithm?
For 2. -> Just make sure nums1 is always small. Just swap the arrays by calling the function like intersect(nums2, nums1)
For 3. -> Load up all nums1 in to a hashmap and then get chunks of num2 and then proceed further as expected.

Intersection of Two Arrays II

Solution
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000


Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

func intersection(nums1, nums2 []int) []int {
	i := 0
	j := 0
	var arr []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			arr = append(arr, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return arr
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	return intersection(nums2, nums1)
}