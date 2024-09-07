package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func evaluateTestFileR(dir string) (bool, error) {
    entries, err := os.ReadDir(dir)
    if err != nil {
        return false, err
    }

    passed := false
    for _, entry := range entries {
        if !entry.IsDir() && 
        strings.HasPrefix(entry.Name(), "turl_") && 
        filepath.Ext(entry.Name()) == ".json" {
            passed = evaluateTestFile(filepath.Join(dir, entry.Name()))
        } else if entry.IsDir() {
            passed, err = evaluateTestFileR(filepath.Join(dir, entry.Name()))
        }
    }

    return passed, nil
}
