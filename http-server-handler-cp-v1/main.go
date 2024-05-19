package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		t := GetTime()
		if t == "" {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

		writer.Write([]byte(t))
	}
}

func GetTime() string {
	t := time.Now()
	return fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
