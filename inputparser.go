package main

import (
	"encoding/json"
	"strings"
)

func getMethod(method string) string {
    valid := []string{
        "GET",
        "POST",
        "PUT",
        "DELETE",
        "PATCH",
    }

    for _, validm := range valid {
        if strings.ToUpper(method) == validm {
            return validm
        }
    }

    return ""
}

func encodeAnyToByte(inputData interface{}) ([]byte, error) {

    switch val := inputData.(type) {
    case string:
        return []byte(val), nil
    case []byte:
        return val, nil
    case nil:
        return nil, nil
    default:
        data, err := json.Marshal(val)
        if err != nil {
            return nil, err
        }
        return data, nil
    }
}
