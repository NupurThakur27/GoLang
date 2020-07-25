package main
import (
	"fmt"
	"sort"
)

func DemoSlice()  {
	count := [10]int{1,2,3,4,5,6,7,8,9,10}
	var slice1 []int = count[0:5]
	var slice2 []int = count[4:10]
	fmt.Println("Sliced have of the count array",slice1)
	slice2[0]=500
	fmt.Println(slice1,slice2)
	fmt.Println(count)
	fmt.Println("Length of slice1 =",len(slice1),"and","capacity of slice1 =",cap(slice1))
}
func Make_slice() {
	arr := make([]int,5)
	//arr := make([]int,0,5) 0 defines the len and 5 defines the capacity
	fmt.Println("Length of the make slice",arr,"is",len(arr),cap(arr))
}

func MultiD_Slice() {
	s1 := [][]int{{1,2},{3,4},{5,6},{7,8}}
	fmt.Println("Multi Dimension Slice or SLice of Slices",s1)
}
func Sort_Slice() {
	s1 := []string{"GoLang","Python","Java","C++","JavaScript","Ruby"}
	fmt.Println("Slice before sorting:",s1)
	sort.Strings(s1) //sort.Ints(s); For int values
	fmt.Println("Slice after sorting",s1)

}