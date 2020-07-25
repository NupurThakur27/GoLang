//a channel is a medium through which a goroutine
//communicates with another goroutine
//and this communication is lock-free
package main
import "fmt"
//Like maps and slices, channels must be created before use

func Demo_Chan() {
	chnl := make(chan string, 5)
	chnl <- "Nupur"
	chnl <- "Thakur"
	chnl <- "GoLang"
	chnl <- "Python"
	//chnl <- "Java"
	//Overflow test
	//chnl <- "C++"
	close(chnl) //Not Necessary
	fmt.Println("Length of the Channel is",len(chnl))
	fmt.Println("Capacity of the Channel is",cap(chnl))
	fmt.Println("Demo Receiving data from Channel",<-chnl)
	for i := range chnl {
		fmt.Println(i)
	}
}
