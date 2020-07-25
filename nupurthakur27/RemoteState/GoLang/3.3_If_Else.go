package main
import "fmt"
func Voting(age int) {
if age < 18 {
	fmt.Println("Grow Up, Kid")
} else {
	fmt.Println("Congratlation You are eligible to vote.")
}

}