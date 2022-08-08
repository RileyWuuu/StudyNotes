package main

import (
	"fmt"
	"sync"
)

func findPrime(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num == 1 {
		return
	} else if num == 2 {
		fmt.Println(num)
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return
			}
		}
		fmt.Println(num)
	}
}

func main() {
	num := 300000
	wg := new(sync.WaitGroup)
	wg.Add(num)
	for i := 1; i <= num; i++ {
		go findPrime(i, wg)
	}
	wg.Wait()
}
