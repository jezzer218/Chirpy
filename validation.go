package main

import (
	"io"
	"net/http"
)

func handleValidation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	type requestBody struct {
		Body string `json:"body"`
	}
	type responseBodySuccess struct {
		Valid bool `json:"valid"`
	}
	type responsBodyError struct {
		Err string `json:"error"`
	}

	dat, err := io.ReadAll(r.Body)
	if err != nil {
		responsBodyError(w, 500, err)
		return
	}

}
