// Should define this one first in every file
package main

import (
	"fmt"
)

// Should only be one main function in the entire package
// The entry points
func main() {
	fmt.Println("This is hello.")
	maps := map[string]int{"A": 1, "B": 2}

	a, prsnt := maps["B"]

	if prsnt {
		hidden := a
		fmt.Println(hidden)
	}

	stringa := "Abcda"

	for i, v := range stringa {
		fmt.Println(i, string(v))
	}
}
