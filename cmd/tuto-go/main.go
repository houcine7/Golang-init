package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

	func main(){
		fmt.Println("Hello my go program is starting ...");
		// var myName string;
		// fmt.Println("Enter your name:");
		//fmt.Scanln(&myName);
		// var result string = greeting(myName);
		// fmt.Println(result)
		
		var num int =10
		var num1 uint16 =16666
		fmt.Println(num,num1)
		fmt.Println("Working with floats")


		var myFloat float32 = 3132.9
		var sum = num + int(myFloat)
		fmt.Println(sum)

		/*
			STRINGS 
		*/
		
		var myString string =`Hello Y`
		var greetingMessage = myString + " " + "I am here";
		fmt.Println("Bytes count : " , len(myString)) // the length is the number of bytes of the string 
		// get the number of characters
		fmt.Println("Character Count: ", utf8.RuneCountInString(myString));
		fmt.Println(myString[0]) // Random access to string elements
		fmt.Println(greetingMessage)

		/*Runes */
		fmt.Println("----------- Runes --------------")
		var firstRune rune =rune(127)
		fmt.Println(string(firstRune))
		

		var v1,v2 string="hello world", "hehe"
		fmt.Println(v1,v2)



		var a,b,err= intDivision(10,2);
		if(err!=nil){
			fmt.Println(err)
		}else if(b ==0){
			fmt.Printf("The result of int division of 10/6 : %v" , a)
		}else{
			fmt.Printf("This the result of int division of 10/6 : %v and the reminder is %v",a ,b)
		}


		//we can perform a switch on a variable
		switch {
			case err !=nil:
				fmt.Println(err)
			case b==0:
				fmt.Printf("The result of int division of 10/6 : %v" , a)
			default:
				fmt.Printf("This the result of int division of 10/6 : %v and the reminder is %v",a ,b)
		}	

	}


	// this my throw a runtime error 
	func intDivision(a int, b int) (int, int, error){
		if(b==0){

			return 0,0, errors.New("Can't divide by zero")
		}
		return a/b , a%b, nil
	}



