package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

const (
	jsonContentType = "application/json"
	textContentType = "application/text"
)

func HttpReturn(w http.ResponseWriter, r *http.Request, statusCode int, contentType string,
	body interface{}) {
	log.Printf("%s %s", r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)

	switch contentType {
	case textContentType:
		output, ok := body.(string)
		if !ok {
			errMsg := fmt.Sprintf("got data type %T, 'string' is required", body)
			err := errors.New(errMsg)
			HttpReturnErr(w, r, err)
		}
		fmt.Fprintf(w, output)
	case jsonContentType:
		json.NewEncoder(w).Encode(body)
	}
}

func HttpReturnErr(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%s %s", r.Method, r.URL)
	errMsg := strings.ReplaceAll(err.Error(), "\n", " ")
	errMsg = strings.ReplaceAll(errMsg, "\t", " ")

	errMsg = fmt.Sprintf("%s %s", r.Method, r.URL)
	zapLogger.Error(errMsg, zap.String("error", err.Error()))
	http.Error(w, err.Error(), 500)
}
