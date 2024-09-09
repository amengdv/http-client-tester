package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const version = "v1.1.0"

func main() {
    args := os.Args[1:]

    versionFlag := flag.Bool("version", false, "Print the version of turl")

    flag.Parse()

    if *versionFlag {
        fmt.Printf("tURL version %s\n", version)
        os.Exit(0) 
    }

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
