package main

import (
	"fmt"
)

func partition(arr []int, p int, r int) int {
	//We always make this a pivot element
	pivot := arr[r]
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j] <= pivot {
			//increment i once for each element in the lower partition
			i++
			tmp := arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
		}
	}
	arr[r] = arr[i+1]
	arr[i+1] = pivot
	return i + 1
}

//Our loop invariant:
//At the beginning of each iteration of the loop , for any array index k,
//1. Ifp≤k≤i,thenA[k]≤x.
//2. Ifi+1≤k≤j-1,thenA[k]>x.
//3. Ifk=r,thenA[k]=x.

//All of this sorting is in place
// So we're done once these calls return`
func quicksort(arr []int, p int, r int) {
	if p < r {
		q := partition(arr, p, r)
		quicksort(arr, p, q-1)
		quicksort(arr, q+1, r)
	}
}

func main() {
	a := []int{8, 1, 0, 40, 31, 1231, 241, 12, -5}
	quicksort(a, 0, len(a)-1)
	fmt.Printf("Answer: %v", a)
}
