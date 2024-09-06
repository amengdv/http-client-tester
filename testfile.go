package main

import (
	"encoding/json"
	"log"
    "os"
)

func evaluateTestFile(filename string) bool {
    tcs := Tests{}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading from file: %v\n", err)
	}

	err = json.Unmarshal(data, &tcs)
	if err != nil {
		log.Fatalln("Error unmarshal json: check your json formatting")
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
    return passed
}
