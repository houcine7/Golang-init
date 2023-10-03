package main

import (
	"fmt"
	"strings"
)

func main() {
	dummyString := "résémue"
	var indexed = dummyString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	// for i,v :=range dummyString{
	// 	fmt.Println(i,v)
	// }

	// to define a string that is represented by  characters
	// we can cast the string to a runes slice
	str := []rune("yes sir")
	fmt.Printf("%v, %T\n", str, str)

	// to build a string from a slice of runes
	chars := []rune{'a','d','u','m','b','a'}
	str2 := ""
	for i :=range chars{
		str2+=string(chars[i])  // each time here a new string is created and the old one is discarded
	}

	// to avoid this we can use the strings.Builder
	var sb strings.Builder

	for i :=range chars{
		sb.WriteString(string(chars[i]))
	}
	
	finalString := sb.String()


	fmt.Println(finalString)




}