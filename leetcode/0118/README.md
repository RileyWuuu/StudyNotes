# 118 Pascal's Triangle

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it
 

**Example 1:**
```
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```
**Example 2:**
```
Input: numRows = 1
Output: [[1]]
```

**Constraints:**

- 1 <= numRows <= 30

## Solutions:
```
(Runtime:2ms)
func generate(numRows int)[][]int{
    result := make([][]int,numRows)
    for i:=0;i<numRows;i++{
        result[i]=make([]int,i+1)
        result[i][0]=1
        result[i][i]=1
        for x := 1;x<i;x++{
            result[i][x]=result[i-1][x-1]+result[i-1][x]
        }
    }
    return result
}
