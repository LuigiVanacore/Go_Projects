package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	CONN_PORT = ":8080"
)

var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		helloWorld(w, r)
	})

var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's a Post Request"))
	})

var PathVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		w.Write([]byte("Hi " + name))
	})

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	router := mux.NewRouter()
	router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetRequestHandler))).Methods("GET")
	logfile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error file server.log: ", err)
	}
	router.Handle("/post", handlers.LoggingHandler(logfile, PostRequestHandler)).Methods("POST")
	router.Handle("/hello/{name}", handlers.CombinedLoggingHandler(logfile, PathVariableHandler)).Methods("GET")
	err = http.ListenAndServe(CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
