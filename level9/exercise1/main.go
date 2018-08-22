package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("From start of main block")
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("From end of main block")
}

func foo() {
	fmt.Println("From foo function")
	wg.Done()
}

func bar() {
	fmt.Println("From bar function")
	wg.Done()
}
