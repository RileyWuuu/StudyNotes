package main

import "fmt"

type test struct {
	name string
}

func main() {
	list := []test{test{"a"}, test{"b"}, test{"c"}}
	for i, v := range list {
		v.name = "update" + string(i)
	}
	fmt.Println(list)
}
