# 977 Squares of a Sorted Array

Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

**Example 1:**

```
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
```

**Example 2:**

```
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
```

**Constraints:**

- 1 <= nums.length <= 104
- -104 <= nums[i] <= 104
- nums is sorted in non-decreasing order.

<hr/>

## Solutions:

### 1. first get the squares of each value, then put them in non-decreasing order by using "sort"

(Runtime:29 ms)

```
func sortedSquares(nums []int)[]int{
    for x:= range nums{
        nums[x] = nums[x]*nums[x]
    }
    sort.Slice(nums,func(a ,b int)bool{
        return nums[a]< nums[b]
    })
    return nums
}
```
### 2. Use two pointer(head & tail) to sort the order of array

```
func sortedSquares(nums []int)[]int{
    head:=0
    tail:=len(nums)-1
    result := make([]int,len(nums))
    current:=len(nums)-1
    for current>=0{
        if math.Abs(float(nums[head]))>=math.Abs(float(nums[tail])){
            result[current]=int(math.Pow(float(nums[head]),2))
            heat++
        }else{
            result[current]=int(math.Pow(float(nums[tail]),2))
            tail--
        }
        current--
    }
    return result
}
```
