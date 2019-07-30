package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handleCall(w http.ResponseWriter, r *http.Request) {
	statusCode := r.URL.Path
	statusCode = strings.TrimPrefix(statusCode, "/")
	code, err := strconv.Atoi(statusCode)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}

	w.WriteHeader(code)
	_, err = w.Write([]byte(statusCode))
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
