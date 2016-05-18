package main

import "fmt"

func bubble(keys []int) []int {
	for i := range keys {
		for j := i + 1; j < len(keys); j++ {
			// if any element is smaller than the one at the bottom, move it to bottom
			// after 0th iteration, a[0] is smallest element
			if keys[i] > keys[j] {
				tmp := keys[j]
				keys[j] = keys[i]
				keys[i] = tmp
			}
		}
	}
	return keys
}

// Loop Invariant lines 7 -> 13 (not 6)!
//Statement:
//At iteration i, keys[0...i] sorted, any element above i greaer than anything in A[0...i]

//Initialization: true prior to first iteration
// At

//Maintenance: true before iteration, true after iteration

//Termination: at completion, we get property proving correctness

func main() {
	fmt.Println("Bubble sort...")
	var a = []int{1, 5, 3, 2, 0, -5}
	var out = bubble(a)
	fmt.Printf("Answer: %v", out)
}
