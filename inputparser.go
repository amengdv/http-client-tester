package main

import "strings"

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
