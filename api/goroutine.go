package handler

import (
	"fmt"
	"net/http"
  "time"
)


func say(s string, w http.ResponseWriter) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Fprintf(w, s)
        }
}

func Handler(w http.ResponseWriter, r *http.Request) {
	go say("world",w)
  say("hello",w)
}
