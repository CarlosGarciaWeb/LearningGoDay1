package main

import (
	"log"
	"time"
)

// package value variable
var s = "seven"

// You can declare variables and their types but this can become messy
// to avoid this you can create a struct that uses a similar format to the
// creating of a database sql table
var firstName string
var lastName string
var phoneNumber string
var age int
// time is a Go built in function that can help establish data types that store date time information
var birthDate time.Time

// struct, the variables are created with capital letters so they can be accessed
// as a package variable ultimately being avaible to be used in funcs withing the
// project file
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	var s2 = "six"

	// if var s string is declared it will expect a new s value inside the func scope
	// you can revalue the package variable through a := inside the func scope

	log.Println("s is ", s)
	log.Println("s2 is ", s2)

	saySomething("xxx")

	// create user variable with shorthand and add information to the struct
	user := User{
		FirstName: "Carlos",
		LastName: "Garcia",
	}

	otherUser := User{
		FirstName: "Juan",
		LastName: "Ramirez",
	}

	// Print the created user variables, since birthdate was not specified it generates a date time
	// type of value of a standard value
	log.Println("This is initial user", user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
	log.Println("This is the new user", otherUser.FirstName, otherUser.LastName)

}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething is", s)
	return s3, "world"
}