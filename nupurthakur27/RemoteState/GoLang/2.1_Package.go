package main
import (
	"fmt"
	"math/rand" // for generating random numbers
	"time"
)
func pkg_basics() {
	fmt.Println("In this time zone, the number assigned to you is",rand.Intn(10))
	/*This is because by default the seed is always the same, the number 1.
	To actually get a random number, you need to provide a unique seed for your program.
	You really want to not forget seeding, and instead properly seed our pseudo number generator.
	How?
	 */
	/*Use rand.Seed() before calling any math/rand method,
	passing an int64 value.
	You just need to seed once in your program,
	not every time you need a random number
	 */
	rand.Seed(time.Now().UnixNano())
	fmt.Println("We had a trick here, the above number is always same but no more tricks.")
	fmt.Println("The ID for you is",rand.Intn(10))

}
