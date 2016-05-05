package main

import "fmt"

//a and b are sorted slices
func merge(a []int, b []int) []int {
  var i, j int = 0, 0
  var out []int
  for i < len(a) && j < len(b) {
    if a[i] < b[j] {
      out = append(out, a[i])
      i++
    } else {
      out = append(out, b[j])
      j++
    }
  }
  if i == len(a) {
    out = append(out, b[j:]...)
  } else {
    out = append(out, a[i:]...)
  }
  return out
}

func merge_sort(keys []int) []int {
  length := len(keys)
  if length > 1 {
    return merge(merge_sort(keys[:length/2]), merge_sort(keys[length/2:]))
  } else {
    return keys
  }
}

func main() {
  fmt.Printf("Merge sort...")
  var a = []int {1, 2, 0, -5, 1231, 1, 135, -2}
  var o = merge_sort(a)
  fmt.Printf("Answer: %v", o)
}
