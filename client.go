package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func sendRequest(inputMethod, inputURL string) (*http.Response, error) {
    client := http.Client{}

    method := getMethod(inputMethod)
    if method == "" {
        method = "GET"
    }

    url, err := url.Parse(inputURL)
    if err != nil {
        return nil, fmt.Errorf("Error parsing url %w\n", err)
    }

    req := &http.Request{
        Method: method,
        URL: url,
    }

    res, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("Error Making HTTP request %w\n", err)
    }


    return res, nil
}

func sendReqWrapper(inputMethod, inputURL string) {

    res, err := sendRequest(inputMethod, inputURL)
    if err != nil {
        log.Fatalf("%v\n", err)
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("%v\n", err)
    }

    fmt.Println(res)
    fmt.Println(string(data))
}
