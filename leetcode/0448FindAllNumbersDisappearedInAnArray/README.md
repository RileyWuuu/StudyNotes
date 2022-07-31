# 448 Find All Numbers Disappeared in an Array

Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

**Example 1:**

```
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
```

**Example 2:**

```
Input: nums = [1,1]
Output: [2]
```

**Constraints:**

- n == nums.length
- 1 <= n <= 105
- 1 <= nums[i] <= n

## Solutions:

### 1. make a copy to show the count of how many times the values shows up, afterwards, check if any value of the copy shows zero, append to new slice and return

(Runtime: 74ms)

```
func findDisappearedNumbers (nums []int)[]int{
    result := []int{}
    count := make([]int,len(nums))
    for i:=0;i<len(nums);i++{
        count[nums[i]-1]+=1
    }
    for i:=0;i<len(count);i++{
        if count[i]==0{
            result = append(result,i+1)
        }
    }
    return result
}
```
