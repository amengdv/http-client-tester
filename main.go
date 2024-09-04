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

    for _, tc := range tcs.TestCases {
        fmt.Println("Name: ", tc.Name)
        sendReqWrapper(&tc)
    }
}
