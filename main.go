package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
    args := os.Args[1:]

    passed := true
    if len(args) == 1 && isDir(args[0]) {
        pass, err := evaluateTestFileR(args[0])
        passed = pass
        if err != nil {
            log.Fatal(err)
        }
    } else {
        for _, v := range args {
            if filepath.Ext(v) != ".json" {
                continue
            }
            pass := evaluateTestFile(v)
            passed = pass
            break
        }
    }

    if passed {
        fmt.Println("----------------------------------")
        log.Println("PASSED")
        fmt.Println("----------------------------------")
    } else {
        log.Println("FAILED")
        fmt.Println("----------------------------------")
    }
}

func isDir(name string) bool {
    filestate, err := os.Stat(name)
    if err != nil {
        log.Println("Cannot find ", name)
        return false
    }

    if filestate.IsDir() {
        return true
    }
    return false
}
