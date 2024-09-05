package main

import "fmt"

func printReport(testName string, passed bool, expected, actual any) {
	if passed {
		fmt.Println("----------------------------------")
		fmt.Println("PASSED")
		fmt.Println("----------------------------------")
		return
	}

	fmt.Println("----------------------------------")
	fmt.Printf("TEST NAME: %v\n", testName)
	fmt.Println("STATUS: FAILED")
	fmt.Printf("EXPECT: %v\n", expected)
	fmt.Printf("ACTUAL: %v\n", actual)
	fmt.Println("----------------------------------")
}
