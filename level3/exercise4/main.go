package main

import "fmt"

func main() {
	bd := 1994
	for {
		if bd > 2018 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
