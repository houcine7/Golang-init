package main

import "fmt"

func main() {
	var slice []int=[]int{1,2,3,5,4,610}
	var slice1 []float32=[]float32{0,7,1.2,0.5,8.5,610}
	fmt.Println(sumSlice[int](slice))

	fmt.Println(sumSlice[float32](slice1))
}


// lets create a function that sums the numbers of a slice for a given slice of  given type 
// to do that we need to write a func for every type 
// here generics would help to get over this problem
func sumSlice[T int | float32 | float64] (slice  []T) T{
		var sum T
		for _, v := range slice{
				sum+= v 
		}
		return sum
}
