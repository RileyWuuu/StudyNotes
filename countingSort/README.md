# Counting sort





##  Explanation:
1. Create an empty array with the length of input array
2. Insert the value of input array into the empty array as key(so if key repeated the value adds.) Which accomplishes the first step of counting sort(count the number of value)
3. Run the array with two for loops, the first for loop represents the key(1-10), and the second for loop represents the count of that key.
4. In the second for loop, if the value doesn't equal to zero(meaning the key exists in the input array), append it into the input array.

### so the code looks like this:
```
func countSort(arr []int)[]int{
    count := make([]int,len(arr)+1)
    for _,x:=range arr{
        count[x]++
    }
    for i,z:=1,0;i<len(count);i++{
        for k:=count[i];k>0;k--{
            arr[z] = i
            z++
        }
    }
    return arr
}
```