package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handleCall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-headers", "*")

	// Get status code from URL.
	statusCode := r.URL.Path
	statusCode = strings.TrimPrefix(statusCode, "/")
	code, err := strconv.Atoi(statusCode)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	// Write status code header.
	w.WriteHeader(code)

	// If no query params are given, the response body becomes the status code.
	if len(r.URL.Query()) == 0 {
		_, err = w.Write([]byte(statusCode))
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		}
		return
	}

	// Else the body becomes the requested body.

	// Parse GET params into struct that we can return as JSON.
	respJSON := make(map[string]interface{}, len(r.URL.Query()))
	for key, values := range r.URL.Query() {
		respJSON[key] = values[0]
	}

	// Set JSON response body.
	jsonBytes, err := json.Marshal(respJSON)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}

func main() {
	http.HandleFunc("/", handleCall)
	fmt.Println("starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println("stopped server")
}
