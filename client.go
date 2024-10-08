package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func sendRequest(inputMethod, inputURL string, inputBody any, header http.Header) (*http.Response, error) {
	client := http.Client{}

	method, err := getMethod(inputMethod)
    if err != nil {
		return nil, fmt.Errorf("ERROR! %w\n", err)
    }

	url, err := url.Parse(inputURL)
	if err != nil {
		return nil, fmt.Errorf("ERROR! Invalid URL\n")
	}

	bodyByte, err := encodeAnyToByte(inputBody)
	if err != nil {
		return nil, fmt.Errorf("ERROR! Can't Encode Body %w\n", err)
	}

	body := io.NopCloser(bytes.NewReader(bodyByte))

	req := &http.Request{
		Method: method,
		URL:    url,
		Header: header,
		Body:   body,
	}

	res, err := client.Do(req)
	if err != nil {
        return nil, fmt.Errorf("ERROR! Can't make HTTP request: %w\n", err)
	}

	return res, nil
}

func sendReqWrapper(tc *TestCase) (bool, testResult) {

	header := tc.Header
	inputMethod := tc.Method
	inputURL := tc.Url
	inputData := tc.InputData

	if header == nil {
		header = &http.Header{}
	}

	res, err := sendRequest(inputMethod, inputURL, inputData, *header)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

    if tc.ShowBody == nil || *tc.ShowBody == true {
        fmt.Println("----------------------------------")
        fmt.Println(tc.Name)
        fmt.Println("Response Body:", string(data))
        fmt.Println("----------------------------------")
    }

	expected := getExpected(tc)
	actual := getActual(res, data)

    return checkResult(expected, actual, tc.Name)
}

func getHeaderKeys(header http.Header) (keys []headerKey) {
    for k := range header {
        keys = append(keys, headerKey(k))
    }
    
    return
}

func getHeaderValues(header http.Header) (values []headerValue) {
    for _, v := range header {
        for _, j := range v {
            values = append(values, headerValue(j))
        }
    }

    return
}
