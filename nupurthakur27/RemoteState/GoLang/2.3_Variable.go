package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func Boolean() bool {
	var b bool
	b = true
	return b
}

func showvar() {
	var a string = ""
	var e int
	fmt.Println("Let's showcase the default values of Integer and String")
	fmt.Println(a)
	fmt.Println(e) //default value of int is 0
	//var b, c int = 1, 2
	//fmt.Println(b, c)
	var b string = "IamString"
	fmt.Println("To showcase the beauty of := operator")
	d := "This is Go Lang" //Just to showcase the beauty of := operator
	f := b + " : " + d     //String concatination
	fmt.Println(f)
}
