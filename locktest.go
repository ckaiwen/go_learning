package main

import (
	"fmt"
	"sync"
)

func main()  {
	a:=1
	p_a:= &a
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	go func() {
		for i:=1;i< 100000;i++{
			lock.Lock()
			*p_a+= 1
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i:= 1;i< 100000 ;i++{
			lock.Lock()
			*p_a-= 1
			lock.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(*p_a)
}