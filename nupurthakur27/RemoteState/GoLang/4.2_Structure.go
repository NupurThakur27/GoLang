package main
import "fmt"
type Name struct {
	fname string
	lname string
}
var (
	n1 = Name{"Nupur","Thakur"}
	n2 = Name{"Remote","State"}
	n3 = Name{"Go","Lang"}
)
func demostruct() {
	fmt.Println("Let's use structure to print your name")
	fmt.Println(Name{"Nupur","Thakur"})
	n := Name{"Nupur","Thakur"}
	n.fname = "XYZ"
	fmt.Println("Experimenting with Struct Fields",n.fname)

	fmt.Println("Struct Lierals",n1,n2,n3)
}
