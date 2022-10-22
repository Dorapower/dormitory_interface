package web

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Hello World")
	if err != nil {
		panic(err)
	}
}
