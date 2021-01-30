package utils

import (
	"fmt"
	"strings"
)

func TrimReturn(str string) string {
	return strings.TrimSuffix(str, "\n")
}

func ErrorHandler(message string, err error) {
	fmt.Println(message, " with error: ", err.Error())
}