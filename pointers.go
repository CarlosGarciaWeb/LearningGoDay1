package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	// Call function changeThePointerUsed and input myString to modify the value in the
	// value location to a new value "red", must use & to target the variable myString location
	changeThePointerUsed(&myString)
	// Show new value
	log.Println("after func call myString is set to", myString)
}


// Referring to a pointer use * and to a reference variable use an &
// Pointers allow to address information stored in a particular location
// for which the below function takes the parameter pointer (location in RAM)
// and replaces it with a new string value "red" in this case the parameter pointer
// must be of a value that is of data type string


//Func below takes s which is a particular position/address in memory of a stored string value
func changeThePointerUsed(s *string) {
	// print the s location that is being inputted into the function
	log.Println(s)
	// creates a new value of data type string using :=
	newValue := "red"
	// gets pointer location and modifies it to be equal to newValue
	*s = newValue
}