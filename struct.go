package main

import "fmt"

/*also note that anything uppercase is public to the other files,
and thus conventions dictate it should be commented, seen here with Student*/

//Student is a sample class
type Student struct {
	name    string
	grade   int
	friends []string
}

func construct() {
	Henry := Student{
		name:  "Henry",
		grade: 95,
		//remember, static typing, so have to delcare list of strings for friends :)
		friends: []string{"Josh", "Ruben"},
	}

	fmt.Println(Henry.grade)
}
