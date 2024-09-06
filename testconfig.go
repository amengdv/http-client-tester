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
    StatusCodeEqual *int `json:"status_code_equal"`
    JsonBodyEqual *json.RawMessage `json:"json_body_equal"`
    BodyEqual *bodyWhole `json:"body_equal"`
    BodyContains *bodySnippet `json:"body_contains"`
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

type bodySnippet string
type bodyWhole string

func getExpected(tc *TestCase) testValue {
	return map[string]any{
		"statusCode":   tc.StatusCodeEqual,
		"expectedJsonBody": tc.JsonBodyEqual,
        "headerKey": tc.HeaderContainsKey,
        "headerVal": tc.HeaderContainsVal,
        "bodyContains": tc.BodyContains,
        "bodyEqual": tc.BodyEqual,
	}

}

func getActual(res *http.Response, data []byte) testValue {
	return map[string]any{
		"statusCode":   res.StatusCode,
		"expectedJsonBody": data,
        "headerKey":  getHeaderKeys(res.Header),
        "headerVal": getHeaderValues(res.Header),
        "bodyContains": bodySnippet(string(data)),
        "bodyEqual": bodyWhole(string(data)),
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
