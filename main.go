package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1]

	data, err := os.ReadFile(args)
	if err != nil {
		log.Fatalf("Error reading from file: %v\n", err)
	}

	tcs := Tests{}

	err = json.Unmarshal(data, &tcs)
	if err != nil {
		log.Fatalf("Error unmarshal json: %v\n", err)
	}

    passed := true
	for _, tc := range tcs.TestCases {
        pass, res := sendReqWrapper(&tc)
        passed = pass
        if !pass {
            printReport(res.testName, false, res.expectedValue, res.actualValue)
            break
        }
	}

    if passed {
        fmt.Println("------------------------------------")
        log.Println("PASSED")
        fmt.Println("------------------------------------")
    } else {
        log.Println("FAILED")
        fmt.Println("------------------------------------")
    }

}
