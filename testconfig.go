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
	InputData *interface{} `json:"input_data"`

	// Test Case Expected Output Field
	StatusCodeEqual *int             `json:"status_code_equal"`
	HeaderEqual     *http.Header     `json:"header_equal"`
	Expected        *json.RawMessage `json:"expected"`
}

type testValue map[string]interface{}

func getExpected(tc *TestCase) testValue {
	return map[string]interface{}{
		"statusCode":   tc.StatusCodeEqual,
		"header":       tc.HeaderEqual,
		"expectedBody": tc.Expected,
	}

}

func getActual(res *http.Response, data []byte) testValue {
	return map[string]interface{}{
		"statusCode":   res.StatusCode,
		"header":       res.Header,
		"expectedBody": data,
	}

}
