package main
import "fmt"
func reverse() {
	fmt.Println("Defer is really useful for reversal operation")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Are you just going to evident that")
}
