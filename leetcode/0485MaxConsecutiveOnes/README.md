# 485 Max Consecutive Ones

Given a binary array nums, return the maximum number of consecutive 1's in the array.

**Example 1:**

```
Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
```

**Example 2:**

```
Input: nums = [1,0,1,1,0,1]
Output: 2
```

**Constraints:**

- 1 <= nums.length <= 105
- nums[i] is either 0 or 1.

<hr/>

## Solutions:

1. Use a for loop to count the number of consecutive 1s, if meet zero then add the counts into a slice. Afterwards compare each count and return the max number.

```
func findMaxConsecutiveOnes(nums []int) int {
    var result int
    var slice []int
    var max int
    for i:=0;i<len(nums);i++{
        if nums[i]==1{
            result ++
        }else{
            slice = append(slice,result)
            result = 0
        }
        if i == len(nums)-1{
            slice = append(slice,result)
        }
    }
    fmt.Println(slice)
    max = slice[0]
    for x :=0;x<len(slice);x++{
        if slice[x]> max{
            max = slice[x]
        }
    }
    return max
}
```

2. Use a for range to and count the 1s by curr\*n+n. Afterwards decide whether it's the max number and return

```
func FindMaxForRange(nums []int) int {
	res, curr := 0, 0

	for _, n := range nums {
		curr = curr*n + n
		if curr > res {
			res = curr
		}
	}
	return res
}
```
