package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	// create a http handler and start the server
	http.HandleFunc("/", handleGoCompileRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("couldn't start the server", err)
	}
}

func handleGoCompileRequest(w http.ResponseWriter, r *http.Request) {
	// parse the request
	code, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "couldn't read the request body", http.StatusBadRequest)
		return
	}
	// compile the go code
	var myStdout, myStderr bytes.Buffer
	i := interp.New(interp.Options{
		Stdout: &myStdout,
		Stderr: &myStderr,
	})
	_ = i.Use(stdlib.Symbols)
	_, err = i.Eval(string(code))
	// send the response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	responseObject := map[string]string{
		"stdout": myStdout.String(),
		"stderr": myStderr.String(),
	}
	if err != nil {
		responseObject["compilation_error"] = err.Error()
	}
	err = json.NewEncoder(w).Encode(responseObject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
