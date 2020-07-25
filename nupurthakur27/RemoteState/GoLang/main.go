package main
import "fmt"
func main() {
	Hello()
	timeZone()
	pkg_basics()
	fmt.Println("Do you want to know what 2+3 is? It's",add(2,3))
	fmt.Println("It's Boolean time, is it true or false? ",Boolean())
	showvar()
	fmt.Println("Let's play with your name now.")
	fmt.Println(swap("Nupur","Thakur"))
	fmt.Println("As we human have caste,religions to discriminate similarly programming Language have Data types")
	VarType()
	fmt.Println("Wait till 10 to proceed further")
	Basic_loop()
	fmt.Println("Looks like you have to wait more")
	WO_Init_Post()
	fmt.Println("Time to play with for loop, or the only loop in Go")
	pattern1()
	Always()
	//IF-ELSE
	Voting(21)
	DayoftheWeek(7)
	//Defer example
	reverse()
	demoptr()
	demostruct()
	Printarray()
	DemoSlice()
	Make_slice()
	MultiD_Slice()
	Sort_Slice()
	Demo_Map()
	Map_Operations()
	//Map Exercise
	f := fibonacci()
	fmt.Println("Fibonacci Series is:")
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	//Method Example
	book := Book{
		bname: "Tarkash",
		author: "Javed Akthar",
		bno: 1,
		ispn: 1235,
	}
	book.Demo_Method()
	//Interface
	var b Cuboid
	b = values{3,4,5}
	fmt.Println("Surface Area of Box:", b.BoxSurfaceArea())
	fmt.Println("Volume of Box:",b.Volume())

	//GORoutine
	go display_name("Go Nupur Thakur")
	display_name("Nupur Thakur")
	//Channels
	Demo_Chan()
	Demo_select()
	Demo_Mutex()


}
