//Goroutine is lightweight thread
//Goroutine is a function or method which executes independently and simultaneously with other Goroutines
//Every concurrent activity in GoLang is called Goroutine
package main

import (
	"fmt"
	"time"
)

func display_name(name string) {
	for i :=0;i<5;i++ {
		time.Sleep(1*time.Second)
		fmt.Println(name)
	}
}


/*Sleep() method in our program which makes the main Goroutine sleeps
for 1 second in between 1-second the new Goroutine executes,
displays “welcome” on the screen,
and then terminate after 1-second main
Goroutine re-schedule and perform its operation. */