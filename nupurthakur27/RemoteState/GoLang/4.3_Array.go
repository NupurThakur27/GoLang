package main
import "fmt"

func Printarray() {
	//Type1 Demo
	var name [2]string
	name[0] = "Nupur"
	name[1] = "Thakur"
	fmt.Println("Individual Array values",name[0],name[1])
	//Type2 Demo
	fmt.Println("Complete Array",name)
	//Type3 Demo
	count := [10]int{1,2,3,4,5,6,7,8,9,10} //another magic of := operator
	fmt.Println(count)
}
