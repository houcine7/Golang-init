package greeting

import "fmt"

func Greeting(name string) string {
	
	// this is a faster way to declare a variable and assign a value to it
	// with the := operator without using the var keyword
	message := fmt.Sprintf("Hello, %v welcome to my go program", name)

	return message

}