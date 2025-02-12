//Write a programme that takes a student, including name, student ID and age (Alternative: store the "date of birth" --> check the time package) and stores it in a struct.
//Create a print function that allows for formatted printing of the student (via a custom String() method on the Student struct).
//Modify the programme to allow for the storing of multiple students and provide an according print functionality.

package main

import "time"

type student struct {
	name string
	age  int
	id   string
	dob  time.Time
}
