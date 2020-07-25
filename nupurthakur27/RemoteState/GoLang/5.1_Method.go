package main

import "fmt"

type Book struct {
	bname string
	author string
	bno int
	ispn int
}
//Method is a function with *reciever* argument
//func Demo_Function(Book b)
//Methods are just the replacement of class in GoLang
//Method can be eclared on non-struct as well
func (b Book) Demo_Method() {
	fmt.Println("Book Name:",b.bname)
	fmt.Println("Author Name:",b.author)
	fmt.Println("Book Number:",b.bno)
	fmt.Println("ISPN:",b.ispn)
}
