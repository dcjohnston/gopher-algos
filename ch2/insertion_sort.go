package main

import "fmt"

func sort(keys []int) []int{
  for i, key := range keys {
    //The only thing to note here is the right-side looping
    //We do this to make it easier to swap values
    //Since we assume the subarray is sorted,
    //Each iteration of this subloop exchanges two elements,
    //"passing" along the key until it reaches its destination.
    for j := i-1; j >= 0; j-- {
      if keys[j] > key {
        keys[j+1] = keys[j]
        keys[j] = key
      }
    }
  }
  return keys
}

func main() {
  fmt.Printf("Insertion sort...")
  var a = []int {8, 1, 0, 40, 31, 1231, 241, 12, -5}
  var o = sort(a)
  fmt.Printf("Answer: %v", o)
}
