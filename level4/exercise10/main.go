package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["richards_goerge"] = []string{"pizza", "alice", "cats"}

	delete(x, "no_dr")

	for k, v := range x {
		fmt.Println("These are the values for:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
