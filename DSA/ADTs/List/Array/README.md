# Array

1. Data structure that contains a fixed numbers of elements that all have the same type.

## Extending a array
1. There's no way to extend an array by simply adding a block behind the current array. Cuz you don't know whether the program is using other spaces for somethine else.

- First we initialize our array, it would be a perfect fit bewteen other spaces.
![array1](https://user-images.githubusercontent.com/71340325/190886039-38e9f794-2b13-44a7-a5c8-38a9bfd51196.jpg)

- If we attend to add one integer "5" into our array, We request another block of memory that can fit our five arrays.
- We copy what we have from the old array into the new array, and free the old memory.
![array2](https://user-images.githubusercontent.com/71340325/190886628-8be6776e-6eb7-4eed-b527-8b50cf391021.jpg)

- However, by copying elements, the time complexity would be O(n)
  - we create a new array and insert the new element to the back(insert_back)
- Better way => we can increase size by 1 for each insert O(1)
