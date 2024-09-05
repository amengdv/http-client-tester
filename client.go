package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func sendRequest(inputMethod, inputURL string, inputBody interface{}, header http.Header) (*http.Response, error) {
    client := http.Client{}

    method := getMethod(inputMethod)
    if method == "" {
        method = "GET"
    }

    url, err := url.Parse(inputURL)
    if err != nil {
        return nil, fmt.Errorf("Error parsing url %w\n", err)
    }

    bodyByte, err := encodeAnyToByte(inputBody)
    if err != nil {
        return nil, fmt.Errorf("Error parsing input data %w\n", err)
    }

    body := io.NopCloser(bytes.NewReader(bodyByte))

    req := &http.Request{
        Method: method,
        URL: url,
        Header: header,
        Body: body,
    }

    res, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("Error Making HTTP request %w\n", err)
    }


    return res, nil
}

func sendReqWrapper(tc *TestCase) {

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

    expect(res, data, tc)
}

