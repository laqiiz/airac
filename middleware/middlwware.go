package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

func Entry(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return Recover(Log(fn))
}

func Recover(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
		defer func() {
			if err := recover(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
				debug.PrintStack()
				http.Error(w, http.StatusText(500), 500)
			}
		}()
	}
}

func Log(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v[byte]", r.Method, r.URL.Path, r.ContentLength)
		start := time.Now()
		fn(w, r)
		log.Printf("%v %v %v %v[byte]", r.Method, r.URL.Path, time.Since(start), r.ContentLength)
	}
}
