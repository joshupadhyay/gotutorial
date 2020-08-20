package main

import (
	"fmt"
)

func slices() {
	a := [...]int{1, 2, 3, 4, 5, 6}

	c := &a

	c[4] = 23

	fmt.Println(a)
	fmt.Println(c)

	//  we can use a slice instead:

	d := []int{1, 2, 3} //initialize a slice with the empty brackets []

	//slices have a length and capacity arguement, which allows us to use less of the entire array

	f := d

	f[0] = 31

	//slices point directly to the underlying data, doesn't make a copy!
}
