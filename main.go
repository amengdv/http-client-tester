package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    args := os.Args[1:]

    passed := true
    if len(args) == 1 && args[0] == "." {
        fmt.Println("DIR")
    } else {
        for _, v := range args {
            pass := evaluateTestFile(v)
            passed = pass
            if !pass {
                log.Println("FAILED")
                fmt.Println("----------------------------------")
                break
            }
        }
    }

    if passed {
        fmt.Println("----------------------------------")
        log.Println("PASSED")
        fmt.Println("----------------------------------")
    }
}
