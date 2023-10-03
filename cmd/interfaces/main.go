package main

import "fmt"

type animal interface{
	speak() string
}

type dog struct{
	breed string
}

func (d dog) speak() string{
	return "woof"
}

func doSomething (a animal) string{
	return a.speak()
}

func main(){
	pug := dog{breed: "pug" }
	fmt.Println(doSomething(pug))
}