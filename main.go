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
	whatWasSaid := saySomething()

	fmt.Println("the function returned", whatWasSaid)
	

}



// the string notation after the func means that the func will in fact return a string value
// in this case it returns "something"

func saySomething() string {

	return "something"

}