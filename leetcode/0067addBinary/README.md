# 67 Add Binary

Given two binary strings a and b, return their sum as a binary string.

 

**Example 1:**
```
Input: a = "11", b = "1"
Output: "100"
```
**Example 2:**
```
Input: a = "1010", b = "1011"
Output: "10101"
``` 

**Constraints:**
```
1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.
```

## solutions:

### 1.Backward summation
```
func addBinary(a string, b string) string {
    result :=""
    carry :=0
    x,y := len(a)-1,len(b)-1
    for x>=0||y>=0{
        c := carry
        if x >= 0 && string(a[x])=="1"{
            c+=1
        }
        if y >= 0 && string(b[y])=="1"{
            c+=1
        }
        if c%2==1{
            result = "1" + result
        }else{
            result = "0" + result
        }
        if c<2{
            carry = 0
        }else{
            carry =1
        }
        x--
        y--
    }
    if carry !=0{
        result = "1"+result
    }
    return result
}
```