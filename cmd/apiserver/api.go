package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jminds/internal/urlshort"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	shorts, err := urlshort.GetShorts()
	path := r.URL.Path

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if url, ok := shorts[path]; ok {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	return
}

func addShortURLs(w http.ResponseWriter, r *http.Request) {
	var shorts urlshort.Shorts

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&shorts); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := urlshort.SaveShorts(shorts); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
