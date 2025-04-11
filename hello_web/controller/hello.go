package controller

import (
	"fmt"
	"net/http"
)

func Handler() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})

	http.ListenAndServe(":8080", nil)
}
