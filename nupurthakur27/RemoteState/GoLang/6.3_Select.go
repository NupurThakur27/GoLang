//select statement is just like switch statement,
//but in the select statement, case statement refers to communication,
//i.e. sent or receive operation on the channel.
package main

import (
	"fmt"
	"time"
)
func Function1(Channel1 chan string) {
	time.Sleep(2*time.Second)
	Channel1 <- "This is Channel1"
}
func Function2(Channel2 chan string) {
	time.Sleep(2*time.Second)
	Channel2 <- "This is Channel2"
}
func Demo_select(){
	c1 := make(chan string)
	c2 := make(chan string)
	go Function1(c1)
	go Function2(c2)
	select {
	case op1 := <-c1:
		fmt.Println(op1)
		case op2 := <-c2:
			fmt.Println(op2)
	}
}