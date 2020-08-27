package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
	//fmt.Fprintf("Hello, World!")
}

