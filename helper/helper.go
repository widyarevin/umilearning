package helper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIRESPONSE(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	// validasi eror required
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}

func RandomValue() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a unique random integer between 0 and 100
	var randomInt int
	unique := false
	existingValues := make(map[int]bool)
	for !unique {
		randomInt = rand.Intn(101)
		if _, exists := existingValues[randomInt]; !exists {
			unique = true
			existingValues[randomInt] = true
		}
	}

	// Concatenate the integer with a string
	randomString := fmt.Sprintf("USER-%d-%t", randomInt, time.Now())

	return randomString
}
