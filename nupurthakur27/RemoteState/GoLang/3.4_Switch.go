package main
import "fmt"

func DayoftheWeek(no int)  {
	switch no { //Similar to if statement switch statement can also have a execution before the condition
	case 1:
		fmt.Println("Today is Sunday")
	case 2:
		fmt.Println("Today is Monday")
	case 3:
		fmt.Println("Today is Tuesday")
	case 4:
		fmt.Println("Today is Wednesday")
	case 5:
		fmt.Println("Today is Thursday")
	case 6:
		fmt.Println("Today is Friday")
	case 7:
		fmt.Println("Today is Saturday")
	default:
		fmt.Println("Enter the correct day number,Nupur.")
	}
}
