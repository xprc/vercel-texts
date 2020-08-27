package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
	//fmt.Fprintf("Hello, World!")
}

