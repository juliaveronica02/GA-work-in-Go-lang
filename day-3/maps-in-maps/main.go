package main

import "fmt"

func main()  {
	// we can store multiple items in a map as well.
	superHero :=map[string]map[string]string {
		"Superman" : map[string]string{
			"realName" : "Clark Kent",
			"city" : "Metropolis",
		},
		"Batman" : map[string]string{
			"realName" : "Bruce Wayne",
			"city" : "Gotham City",
		},
	}
	// we can output data where the key matches superman.
	if temp, hero := superHero["Superman"]; hero{
		fmt.Println(temp["realName"], temp["city"])
	}
}