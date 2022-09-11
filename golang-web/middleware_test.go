package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(w, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Executed")
		panic("Ups")
	})

	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}

	ErrorHandler := &ErrorHandler{
		Handler: LogMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: ErrorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
