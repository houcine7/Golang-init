package main

import (
	"fmt"
	"unsafe"
)

func dummy(n *int) {
	*n += 1
}

func manipArr(arr *[3]int) {
	(*arr)[0] = 57
}

// this is an advanced use of arrays

func pointerArithmetic(arr []int) int {
	var ptr *int = &arr[0]
	for i := range arr {
		fmt.Println(*ptr, i)
		ptr++
	}
	return *ptr

}

func main() {

	var intPointer *int64 = new(int64)
	*intPointer = 10
	fmt.Printf("ThiS is the pinter address %v points to value %v\n", intPointer, *intPointer)
	intPointer = nil

	fmt.Println("Calling ... the dummy function")
	var n int = 25

	dummy(&n)
	fmt.Println(n)

	//Array
	//arr := [3]int{1, 2, 5}
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var ptr *int = &arr[0]
	for i := 0; i < len(arr); i++ {
		fmt.Println(*ptr)
		ptr++
	}

	ptrToFirstIndex := &arr[0]
	fmt.Printf("the pointer %v, from the arr points to %v in the arr \n", ptrToFirstIndex, *ptrToFirstIndex)

	// Arithmitic Pointers are not supported by go directly
	// ptrToFirstIndex++

	// we can use unsafe package to support it

	ptrToFirstIndex = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptrToFirstIndex)) + uintptr(unsafe.Sizeof(arr[0]))))

	fmt.Printf("the pointer after incrementing by one %v, from the arr points to %v in the arr \n", ptrToFirstIndex, *ptrToFirstIndex)

	// dividing array elements by 2
	fmt.Println(arr)
	manipArr(&arr)
	fmt.Println(arr)

}
