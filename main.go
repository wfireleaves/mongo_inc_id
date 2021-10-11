package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	factory := NewIDFactory()
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			id, err := factory.NextID()
			if err != nil {
				fmt.Println(num, "nextID err,", err.Error())
			}
			fmt.Println(num, ":", id)
		}(i)
	}
	wg.Wait()
}
