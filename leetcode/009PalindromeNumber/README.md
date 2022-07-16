# 009 Palindrome Number

## Solution:

### 1. Reverse the input number (using for loop) and compare

```
func isPalindrome(x int) bool {
    tmp := x
    var rn int
    for tmp> 0{
        rn = rn*10 + tmp%10
        tmp =tmp/10
    }
    return x == rn
}
```
