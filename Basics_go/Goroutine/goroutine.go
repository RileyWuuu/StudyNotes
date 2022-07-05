package main

import (
	"fmt"
	"sync"
)

func main() {

	samples := []string{"cloud", "d2nova", "evox"}
	wg := new(sync.WaitGroup)
	for i, _ := range samples {
		wg.Add(1)
		go getAllvalue(samples[i], wg)
	}
	wg.Wait()
}

func getAllvalue(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
