package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func expect(res *http.Response, data []byte, tc *TestCase) {
    tcExpect, err := encodeAnyToByte(tc.Expected)
    if err != nil {
        log.Println("Error encode tc.Expected")
        return
    }

    expected :=  map[string]interface{}{
        "statusCode": tc.StatusCodeEqual,
        "header": tc.HeaderEqual,
    }

    actual :=  map[string]interface{}{
        "statusCode": res.StatusCode,
        "header": res.Header,
    }


    pass :=  true

    if tc.Expected != nil && !bytes.Equal(tcExpect,  data) {
        fmt.Println("EXPECT: ", string(tcExpect))
        fmt.Println("ACTUAL: ", string(data))
        fmt.Println("TEST FAILED")
        return
    }

    for key, expectedVal := range expected {
        fmt.Println("ON KEY: ", key)
        expectedValReflect := reflect.ValueOf(expectedVal)

        if expectedValReflect.Kind() == reflect.Ptr && expectedValReflect.IsNil() {
            continue
        }

        actualVal := actual[key]

        // fmt.Println("Expected Pointer: ", expectedVal)
        // fmt.Println("Expected Val Reflect: ", expectedValReflect)
        // fmt.Println("Expected Type Reflect: ", reflect.TypeOf(expectedVal))

        copyExpectedVal := extractValueFromPointer(expectedValReflect)

        fmt.Printf("Expect: %v\n", copyExpectedVal)
        fmt.Printf("Actual: %v\n", actualVal)

        switch v := copyExpectedVal.(type) {
        case int:
            if v != actualVal.(int) {
                pass = false
            }
        case http.Header:
            if !reflect.DeepEqual(v, actualVal.(http.Header)) {
                pass = false
            }
        }

        if !pass {
            fmt.Println("TEST FAILED")
            return
        }
    }


    fmt.Println("TEST PASSED")

}

func extractValueFromPointer(val reflect.Value) any {
    if val.Kind() != reflect.Ptr {
        return nil
    }
    return reflect.Indirect(val).Interface()
}
