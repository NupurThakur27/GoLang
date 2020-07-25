package main

import "fmt"
func Basic_loop() {
	for i :=0; i<10; i++ {
		fmt.Println(1 + i)
	}
}
func WO_Init_Post() {
	i := 1
	for i<=10 {      //Wait, Does it look like while loop to you too?
		fmt.Println(i)
		i = i+1
	}
}