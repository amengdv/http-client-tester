package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strings"
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

        // Check result base on type
		switch v := copyExpectedVal.(type) {
		case int:
			if v != actualVal.(int) {
				pass = false
			}
        case headerKey:
            exist := headerContainsKey(v, actualVal.([]headerKey))
            if !exist {
                pass = false
            }
        case headerValue:
            exist := headerContainsValue(v, actualVal.([]headerValue))
            if !exist {
                pass = false
            }
        case bodySnippet:
            exist := isBodyContains(v, actualVal.(bodySnippet))
            if !exist {
                pass = false
            }
        case bodyWhole:
            equal := isBodyEqual(v, actualVal.(bodyWhole))
            if !equal {
                pass = false
            }
		}

		if !pass {
			return false, failTestResult(testName, copyExpectedVal, actualVal)
		}
	}
    return true, successTestResult(testName)
}


// HEADER KEY AND VALUE CHECKER
func headerContainsKey(expect headerKey, actual []headerKey) bool {
    for _, v := range actual {
        if strings.ToLower(string(v)) == strings.ToLower(string(expect)) {
            return true
        }
    }
    return false
}

func headerContainsValue(expect headerValue, actual []headerValue) bool {
    for _, v := range actual {
        if strings.ToLower(string(v)) == strings.ToLower(string(expect)) {
            return true
        }
    }
    return false
}

// BODY CHECKER
func isBodyEqual(expect, actual bodyWhole) bool {
    if string(expect) == string(actual) {
        return true
    }
    return false
}


func isBodyContains(expect, actual bodySnippet) bool {
    if strings.Contains(string(actual), string(expect)) {
        return true
    }

    return false
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
	tcExpect, err := encodeAnyToByte(expected["expectedJsonBody"].(*json.RawMessage))
	if err != nil {
		// Just log the error
		// Can't stop checking other type just because of this error
		log.Println("Error Decoding Expected Body Field")
	}

    actualBody := actual["expectedJsonBody"].([]byte)

	if bodyEqual, err := checkBodyEqual(actualBody, tcExpect); err != nil {
		if err != notDefinedError {
			log.Println("Error encode expected field")
		}
	} else {
		if bodyEqual == false {
            actArgs := string(actualBody)
            expArgs := string(tcExpect)
            return failTestResult(testName, actArgs, expArgs), errors.New("Assert Expected != Actual")
		}
	}

    return successTestResult(testName), nil
}
