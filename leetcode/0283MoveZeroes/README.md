# 283 Move Zeroes

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

 

**Example 1:**
```
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
```
**Example 2:**
```
Input: nums = [0]
Output: [0]
```

**Constraints:**

- 1 <= nums.length <= 104
- -231 <= nums[i] <= 231 - 1

## Solutions:

### 1.Use two pointers and exchange each number when the value != 0
(Runtime: 21ms)
```
func moveZeroes(arr []int){
    x :=0
    for i:=0;i<len(arr);i++{
        if nums[i]!=0{
            nums[i],nums[x] = nums[x],nums[i]
            x++
        }
    }
}
```
