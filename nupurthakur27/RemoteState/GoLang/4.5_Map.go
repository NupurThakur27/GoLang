package main

import "fmt"

type Name_map struct {
	fname, lname string
}
var m map[string]Name_map
func Demo_Map(){
	m=make(map[string]Name_map)
	m["My name is"]=Name_map{"Nupur","Thakur"}
	fmt.Println("Value of My name is",m["My name is"])
}