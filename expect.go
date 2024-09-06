package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
)

var notDefinedError error = errors.New("User does not defined the field")


func checkResult(expected, actual testValue, testName string) (bool, testResult) {

	// Check for response body first
	checkBodyResult, err := checkBodyWrapper(expected, actual, testName)
	if err != nil {
		return false, checkBodyResult 	
    }

	// Check for other type
	pass := true
	for key, expectedVal := range expected {
		expectedValReflect := reflect.ValueOf(expectedVal)

		if expectedValReflect.Kind() == reflect.Ptr && expectedValReflect.IsNil() {
			continue
		}

		actualVal := actual[key]
		copyExpectedVal := extractValueFromPointer(expectedValReflect)

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
			return false, failTestResult(testName, copyExpectedVal, actualVal)
		}
	}


    return true, successTestResult(testName)
}

func checkBodyEqual(actual []byte, expect []byte) (bool, error) {
	if string(expect) == "null" {
		return false, notDefinedError
	}

	if !bytes.Equal(actual, expect) {
		return false, nil
	}

	return true, nil
}

func checkBodyWrapper(expected, actual testValue, testName string) (testResult, error) {
	tcExpect, err := encodeAnyToByte(expected["expectedBody"].(*json.RawMessage))
	if err != nil {
		// Just log the error
		// Can't stop checking other type just because of this error
		log.Println("Error Decoding Expected Body Field")
	}

	if bodyEqual, err := checkBodyEqual(actual["expectedBody"].([]byte), tcExpect); err != nil {
		if err != notDefinedError {
			log.Println("Error encode expected field")
		}
	} else {
		actualBody := actual["expectedBody"].([]byte)
		if bodyEqual == false {
            actArgs := string(actualBody)
            expArgs := string(tcExpect)
            return failTestResult(testName, actArgs, expArgs), errors.New("Assert Expected != Actual")
		}
	}

    return successTestResult(testName), nil
}
