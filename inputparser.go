package main

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

func getMethod(method string) (string, error) {

    if len(method) == 0 {
        return "GET", nil
    }

	valid := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
	}

	for _, validm := range valid {
		if strings.ToUpper(method) == validm {
			return validm, nil
		}
	}

	return "", errors.New("Unsupported HTTP Method")
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

func extractValueFromPointer(val reflect.Value) any {
	if val.Kind() != reflect.Ptr {
		return nil
	}
	return reflect.Indirect(val).Interface()
}
