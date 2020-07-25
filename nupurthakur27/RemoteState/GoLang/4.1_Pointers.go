package main
import "fmt"
//Unlike C, Go has no pointer arithmetic.
func demoptr() {
	i,j := 2,3
	p := &i
	fmt.Println("Here i am printer a pointer value",*p)
	*p = j
	fmt.Println("The actual value after changing the pointer value",i)

}
