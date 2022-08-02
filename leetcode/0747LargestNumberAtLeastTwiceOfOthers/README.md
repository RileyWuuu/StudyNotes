# 747 Largest Number At Least Twice of Others

You are given an integer array nums where the largest integer is unique.

Determine whether the largest element in the array is at least twice as much as every other number in the array. If it is, return the index of the largest element, or return -1 otherwise.

**Example 1:**

```
Input: nums = [3,6,1,0]
Output: 1
Explanation: 6 is the largest integer.
For every other number in the array x, 6 is at least twice as big as x.
The index of value 6 is 1, so we return 1.
```

**Example 2:**

```
Input: nums = [1,2,3,4]
Output: -1
Explanation: 4 is less than twice the value of 3, so we return -1.
```

**Constraints:**

- 2 <= nums.length <= 50
- 0 <= nums[i] <= 100
- The largest element in nums is unique.

## Solutions:

### run through the array to get the index and value of maximum slice, compare with each value and return the index.

```
func dominantIndex(nums[]int)int{
    max,idx:=0,0
    for i,x:=range nums{
        if max<x{
            max = x
            idx=i
        }
    }
    for _,x:=range nums{
        if x!=max && x*2>=max{
            return -1
        }
    }
    return idx
}
```
