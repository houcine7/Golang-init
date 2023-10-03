package main

import "fmt"


/*
	Why we use struct?
		- to group related data together
		- to create custom types
		- to create methods on types
		- to implement interfaces
		- to create anonymous types
		- to create embedded types
		- to create nested types
		- to create exported types
		- to create unexported types

*/

type user struct{
	name string
	age int
	isAuthorized bool
	address address // we ca say=> address // without the field name if the type is the same as the field name
}

// we can add methods to struct
// are functions that are attached to a struct
func (u user) isAuthenticated() bool{
	return u.isAuthorized
}

type address struct{
	num int 
	city string
	street string
}


/*
	Reasons why to use interfaces
*/


type dog struct{
	breed string
}

func (d dog) speak(){
	fmt.Println("woof")
}

type cat struct{
	eyeColor string
}

func (c cat) speak(){
	fmt.Println("meow")
}

// this is the problem we have to create a function for each type
// we can use interfaces to solve this problem


func main(){
	var u user = user{name:"Houcine", age:22, isAuthorized: false, address: address{num: 22, city: "Tunis", street: "Rue 22"}};
	var u2 user
	u2.name="hassan"
	u2.age= 22
	u2.isAuthorized= true
	u2.address.city="Tunis"
	u2.address.num= 22
	u2.address.street= "Rue 22"
	fmt.Println(u2,u)
	// we can use anonymous struct they are 
	//useful when we want to create a struct that we will use only once
	var doc  = struct{
				name string
				pages int
			}{"Golang", 22}

	fmt.Println(doc)
	
}