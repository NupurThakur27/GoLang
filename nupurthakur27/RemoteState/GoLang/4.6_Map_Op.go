package main
import "fmt"
func Map_Operations() {
	m :=make(map[int]string)
	m[1]="Python"//Insertion
	fmt.Println("The value of Map after insertion is",m[1])
	m[1]="GoLang"//Updation
	fmt.Println("The value of Map after updating is",m[1])
	delete(m,1)//Deletion
	fmt.Println("The value of Map after deletion is",m[1])
	value, ok := m[1]
	fmt.Println("The value of Map is ",value,"Is it present?",ok)

}
