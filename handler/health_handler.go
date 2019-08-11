package handler

import (
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "ok")
}
