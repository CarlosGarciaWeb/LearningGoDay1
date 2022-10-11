package main

import "fmt"
// Print Hello, world! to the cli
func main() {
	fmt.Println("Hello, world!")


	// create variable of string named whatTosay and call it with fmt.Println
	var whatToSay string
	// create variable int depending on bit computer. Int can be used as standar.

	var i int

	whatToSay = "Let's say hello again, world!"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("is is set to", i)


	// using the := creates a variable with the same data type of the input, in this case
	// the input is the return output from the function saySomething which returns a 
	// string with value "something"
	whatWasSaid, otherThingSaid := saySomething()

	fmt.Println("the function returned", whatWasSaid, otherThingSaid)
	

}



// the string notation after the func means that the func will in fact return a string value
// in this case it returns "something", a second denoted string means the function will
// return a second output of string with value "else"
// however, this new output must be used or it will generate an error

func saySomething() (string, string) {

	return "something", "else"

}