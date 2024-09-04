package main

type Tests struct {
    TestCases []TestCase `json:"tests"`
}

type TestCase struct {
    Name string `json:"name"`
    Method string `json:"method"`
    Url string `json:"url"`
    StatusCodeEqual *int `json:"status_code_equal"`
    ContentTypeEqual *string `json:"content_type_equal"`
}
