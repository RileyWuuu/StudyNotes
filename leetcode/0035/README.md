# 35. Search Insert Position

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.

**Example 1:**
```
Input: nums = [1,3,5,6], target = 5
Output: 2
```
**Example 2:**
```
Input: nums = [1,3,5,6], target = 2
Output: 1
```
**Example 3:**
```
Input: nums = [1,3,5,6], target = 7
Output: 4
```

**Constraints:**

* 1 <= nums.length <= 104
* -104 <= nums[i] <= 104
* nums contains distinct values sorted in ascending order.
* -104 <= target <= 104

<hr/>

## Solution:

### 1.Use a simple for loop, if neither of them match target then return the length of slice
(Runtime: 10 ms)
```
func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
```
}
### 2.Use sort.Search to find the target number, sort returns the index of the value
(Runtime: 2 ms)
```
func searchInsert(nums []int, target int) int {
    i := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= target
    })
    return i
}
```
