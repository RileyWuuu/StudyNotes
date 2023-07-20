# 001 TwoSums

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- 2 <= nums.length <= 104
- -109 <= nums[i] <= 109
- -109 <= target <= 109
- Only one valid answer exists.

## Solutions:

### 1. Calculate and find whether the value of (target - x) exists in the input slice and replace the x(value) to i(key) and return
Time Complexity:O(n)
```
func TwoSums(nums []int, target int) []int {
	intMap := make(map[int]int)
	for i, x := range nums {
		if v, ok := intMap[target-x]; ok {
			return []int{v, i}
		}
		intMap[x] = i
	}
	return []int{}

}
```
### 2. Solve in a straight forward (but inefficent) way with two for loop
Time Complexity:O(n^2)

```
func twoSums(nums []int)[]int{
	result := make([]int,2)
	for i:=0;i<len(nums);i++{
		for x:=i+1;x<len(nums);x++{
			if nums[i]+nums[x]==target{
				result[0]=i
				result[1]=x
			}
		}
	}
	return result
}
```
```
