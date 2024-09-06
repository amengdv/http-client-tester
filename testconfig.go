package main

import (
	"encoding/json"
	"net/http"
)

type Tests struct {
	TestCases []TestCase `json:"tests"`
}

type TestCase struct {
	// Name of the test case
	Name string `json:"name"`

	// Test case input
	Method    string       `json:"method"`
	Url       string       `json:"url"`
	Header    *http.Header `json:"header"`
	InputData *any `json:"input_data"`

	// Test Case Expected Output Field
	StatusCodeEqual *int             `json:"status_code_equal"`
	BodyEqual        *json.RawMessage `json:"body_equal"`
    HeaderContainsKey *headerKey `json:"header_contain_key"`
    HeaderContainsVal *headerValue `json:"header_contain_value"`
}

type testResult struct {
    testName string
    testStatus string
    expectedValue any
    actualValue any
}

type testValue map[string]any

type headerKey string
type headerValue string

func getExpected(tc *TestCase) testValue {
	return map[string]any{
		"statusCode":   tc.StatusCodeEqual,
		"expectedBody": tc.BodyEqual,
        "headerKey": tc.HeaderContainsKey,
        "headerVal": tc.HeaderContainsVal,
	}

}

func getActual(res *http.Response, data []byte) testValue {
	return map[string]any{
		"statusCode":   res.StatusCode,
		"expectedBody": data,
        "headerKey":  getHeaderKeys(res.Header),
        "headerVal": getHeaderValues(res.Header),
	}
}

func failTestResult(testName string, expectedVal, actualVal any) testResult {
    return testResult{
        testName: testName,
        testStatus: "FAIL",
        expectedValue: expectedVal,
        actualValue: actualVal,
    }
}

func successTestResult(testName string) testResult {
    return testResult{
        testName: testName,
        testStatus: "SUCCESS",
        expectedValue: nil,
        actualValue: nil,
    }
}
