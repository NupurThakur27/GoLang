package main

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a, b int
	return func() int {
		a, b = b, a + b
		if b == 0 {
			b = 1
		}
		return a
	}
}

