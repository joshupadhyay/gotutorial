package main

import "fmt"

func studentMaps() {

	grades := map[string]int{
		"Bryce":     84,
		"Timothy":   65,
		"Frederick": 95,
		"Chumlee":   98,
		"Harry":     82,
	}

	//correct key, all ok
	fmt.Println(grades["Bryce"])

	//incorrect key, no error, returns 0! Stops prgm from breaking, but could be bad
	fmt.Println(grades["Brye"])

	//enter the ok statement!

	_, ok := grades["Brye"]

	fmt.Println(ok)

	//now we can see pgrm returns false, indicating key is not in the map

	//also keep in mind, maps are like slices, point directly to reference

}
