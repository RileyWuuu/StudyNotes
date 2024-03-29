# 1346 Check If N and Its Double Exist

Given an array arr of integers, check if there exists two integers N and M such that N is the double of M ( i.e. N = 2 \* M).

More formally check if there exists two indices i and j such that :

- i != j
- 0 <= i, j < arr.length
- arr[i] == 2 \* arr[j]

**Example 1:**

```
Input: arr = [10,2,5,3]
Output: true
Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.
```

**Example 2:**

```
Input: arr = [7,1,14,11]
Output: true
Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
```

**Example 3:**

```
Input: arr = [3,1,7,11]
Output: false
Explanation: In this case does not exist N and M, such that N = 2 * M.
```

**Constraints:**

- 2 <= arr.length <= 500
- -10^3 <= arr[i] <= 10^3

## Solutions:

### 1. put the value of slice into map as key, check if value*2 exist in the slice,(also filter out 0 since 0*2=0)

(Runtime: 5ms)

```
func checkIfExist(arr []int)bool{
    check := make(map[int]int)
    for x:=range arr{
        check[arr[x]] = x
    }
    for x := range arr{
        if count,ok:=check[arr[x]*2];ok && count != x{
            return true
        }
    }
    return false
}
```

### 2. use a for loop in a for loop to compare each value in slice

(Runtime: 4ms)

```
func checkIfExistII(arr []int)bool{
    ln := len(arr)
    for i:=o;i<ln;i++{
        for x:=i+1;x<ln;x++{
            if arr[i]==arr[x]*2||arr[x]==arr[i]*2{
                return true
            }
        }
    }
    return false
}
```
