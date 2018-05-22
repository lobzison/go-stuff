package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    sum := 0
	prev_val := 0
	prev_val2 := 0
	return func() int {
	    sum = prev_val + prev_val2
		prev_val2 = prev_val
		prev_val = sum
		if prev_val == 0 {
		    prev_val2++
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
