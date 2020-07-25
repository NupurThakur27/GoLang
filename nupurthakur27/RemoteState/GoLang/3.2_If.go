package main

import (
	"fmt"
)

func Always() {
	if v := 5; v == 5 {    //This if statement contains a execution statement before execution
		fmt.Println("This condition is always true")
	}
}

