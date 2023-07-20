# 7/18 Arrays 101

## 1. What is an Array?

- a collection of items.
- could be int, string...etc.
- items are stores in contiguous memory locations.

### 2. Creating an Array

- Arrays can hold up to N items.(N is decided by you)
- You need to specify the type of item that goes into the Array.

- **Go:**

```
func createArray(){
    var arr [1]int
    var arr2 [1]string
}
```

<hr/>
* If you declared a length 1000000 array, the computer will reserve memory to hold 1000000 values.

## 3. Writing into an Array

### Declare the index we want to write in

- **Go:**

```
func writeIntoArray(){
    var arr [3]string
    arr[2] = "STOP"
}
```

- Arrays can be overwritten, for example:

```
func overWrite(){
    var arr [3]string
    arr[2] = "STOP"
    arr[2] = "COOL"
    fmt.Println(arr[2])
}
```

- the result would be "COOL"

### With a loop

- **Go:**

```
func writeWithloop() {
	var arr [3]int
	for i := 0; i < len(arr); i++ {
		input := (i + 1) * (i + 1)
		arr[i] = input
	}
	fmt.Println(arr)
}
```

- the result would be [1,4,5]

## 4. Reading from an Array

### reading with index

- If we try to access an index with no value, we will get ""(for string array), 0 (for int array), and false (for bool array)

### reading with loop

- **Go:**

```
func readWithloop(data [10]int) {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}
```

**OR**

```
func readWithForRange(data[10]int){
    for _,i := range data{
        fmt.Println(i)
    }
}
```

## 5. Insert, Delete and Search

### Insert

- To insert an element to the start of an Array, we need to shift all elements backwards(to the right)

**Go:**

```
func shiftElement(data [10]int) [10]int {

	for i := 5; i >= 0; i-- {
		data[i+1] = data[i]
	}
	return data
}

```

- To insert at any given index, we need to shift all elements from that index onwards to the right.

**Go:**

```
func insertAnyIndex(data [10]int) [10]int {
	for i := 6; i >= 3; i-- {
		data[i+1] = data[i]
	}
	data[3] = 777
	return data
}
```

### Delete

- Delete at the end of an Array also known as a queue.

**Go:**

```
func deleteFirstIndex(data [10]int) [10]int {
	for i := 1; i < len(data); i++ {
		data[i-1] = data[i]
	}
	return data
}
```

- Delete from anywhere

**Go:**

```
func deleteAnyIndex(data [10]int) [10]int {
	for i := 3; i < len(data); i++ {
		data[i-1] = data[i]
	}
	return data
}
```

### Search

- Linear search: time complexity/O(N)
  - if the index is unknown, we can check every element in the array unitl we find the target.
- Binary search:
  - if the elements are in sorted order, we can use binary search which we repeatedly look at the middle element, and determine whether the element we're looking is on the left or right side. For each time we halve the number of elements.

## 6.In-Place Operations

- In-place Array helps reduce space complexity

  - E.g. O(n) -> O(1)

- When we iterate over the Array in two different places at the same time, we use **two pointer technique**
  - readPointer
  - writePointer
