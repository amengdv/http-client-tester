package main

import "testing"

func TestInputDataParse(t *testing.T) {
    tests := []struct{
        name string
        input interface{}
        expected string
    }{
        {
            name: "test with struct",
            input: struct{
                Email string `json:"email"`
                Password string `json:"password"`
            }{
                Email: "saul@bettercall.com",
                Password: "123456",
            },
            expected: `{"email":"saul@bettercall.com","password":"123456"}`,
        },
        {
            name: "test with integer",
            input: 200,
            expected: "200",
        },
        {
            name: "test with boolean",
            input: false,
            expected: "false",
        },
        {
            name: "test with struct",
            input: struct{
                ID int `json:"id"`
                Email string `json:"email"`
                Yes bool `json:"yes"`
            }{
                ID: 1,
                Email: "saul@bettercall.com",
                Yes: false,
            },
            expected: `{"id":1,"email":"saul@bettercall.com","yes":false}`,
        },
    }

    for i, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            body, err := encodeAnyToByte(tc.input)
            if err != nil {
                t.Errorf("Test %v FAIL, Error: %v\n", i, err)
            }

            actual := string(body)
            if actual != tc.expected {
                t.Errorf("Test %v FAIL, Expect %v Got %v\n", i, tc.expected, actual)
            }
        })
    }
}
