package main

import (
	"fmt"
	"time"
)


func arrTest(){
	var arr [3]int32; // init the array with default values
	fmt.Println(arr);
	arr[0] = 1;
	arr[1] = 8;
	arr[2] = 3
	fmt.Println(arr[2]);
	
	// we can get a slice of the array
	var slice []int32 = arr[0:2]; // last element is excluded
	fmt.Println(slice);

	// Go will require 4 * 3 bytes of memory for this array 
	// this memory is continuos and allocated  on the stack
	/*fmt.Printf("The memory address of %v  is %v          ", 0,&arr[0]);
	fmt.Printf("The memory address of %v  is %v     ", 1,&arr[1]);
	fmt.Printf("The memory address of %v  is %v       ", 2,&arr[2]);*/ //As we can see the memory address of the array elements are contiguous address2 =address1+4bytes


	// we can declare and init an array 
	var arr4 =[2]int32{1,2};
	fmt.Println(arr4);
	//OR
	arr5, arr6:= [2]int32{1,2},[...]int32{5,4,6};
	fmt.Println(arr5,arr6);
}

func sliceTest(){
	// slices 
	var mySlice []int32 = []int32{1,4,5,7} ; // 
	fmt.Println(mySlice, "the size", len(mySlice), "capacity",cap(mySlice)); 
	// the capacity is the number of elements that the slice can hold before it needs to be resized 
	// the size is the number of elements in the slice
	//the capacity will equal the size of the underlying array if we don't specify it
	mySlice = append(mySlice, 8,32)
	fmt.Println(mySlice,"the size", len(mySlice), "capacity",cap(mySlice))
	// when we want to append a new element to the array and the capacity is not enough,
	// Go will create a new array with double the capacity(len(arr)*2) and 
	//copy the elements of the old array to the new one


	// we create a slice from an array
	var arr7 = [5]int32{1,2,3,4,5};
	var slice2 = arr7[0:3];
	slice2[0] = 100;

	//with make method
	slice3 := make([]int32 , 3 , 10);
	fmt.Println(slice3)
}

func iterations(){
	// Iterate over an array
	arrToIterate := [...]int32{2,45,4,40,1,2,6,5,2,6,4}
	for index,num :=range arrToIterate{
		if(index % 2==0) {
			continue;
		}
		fmt.Println(num)
	}

	//Iterate over a slice 
	sliceToIterate := arrToIterate[:]
	for index,num :=range sliceToIterate{
		if(index % 2==0) {
			continue;
		}
		fmt.Println(num)
	}

	for index:=0;index<len(sliceToIterate);index++{
		fmt.Println(sliceToIterate[index])
	}

	// in go there is no while loop
	// we can use for loop to simulate while loop

	var i int=0;
	for i<10{
		fmt.Println(i)
		i++;
	}
	var j int =0;
	for{
		if(j==9){
			break;
		}
		fmt.Println(j);
		j++;
	}
}

func slicePerformance(){
	n:=1000000;
	//slice 1 without capacity
	var slice1 = []int{}
	//slice 2 with capacity
	var slice2 = make([]int,0,n/3);	

	fmt.Println("slice 1 took ", loopDuration(slice1,n))
	fmt.Println("Slice 2 took", loopDuration(slice2,n))
}


func loopDuration(slice []int , n int) time.Duration{
	var t0 =time.Now();
	for len(slice)<n{
		slice = append(slice,1);
	}
	return time.Since(t0);
}



func main(){
	slicePerformance()
}