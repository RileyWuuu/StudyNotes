# LinkedList
```
type linkedList struct{
    head *Node
    tail *Node
    length int
}
```
- Composed of multiple nodes
- Each node contains a value and pointer to the next list
- Nodes are not arranged sequentially in memory
- Node: <br>
![node](https://user-images.githubusercontent.com/71340325/188661020-949841cf-6ba9-4309-9b3e-1474e1e201fa.PNG)

```
type Node struct{
    value int
    next *Node
}
```

## Advantages:
- Length can easily change and does not need to be known in advance. (Dynamically allocates nodes)
- Removing / Adding nodes at front in O(1)
