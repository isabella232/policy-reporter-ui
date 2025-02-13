package api

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/kyverno/policy-reporter-ui/pkg/config"
	"github.com/kyverno/policy-reporter-ui/pkg/report"
)

func PushResultHandler(store *report.ResultStore) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		var result report.Result

		err := json.NewDecoder(req.Body).Decode(&result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{ "message": "%s" }`, html.EscapeString(err.Error()))
		}

		store.Add(result)
	}
}

func ResultHandler(store *report.ResultStore) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		err := json.NewEncoder(w).Encode(store.List())
		if err != nil {
			fmt.Fprintf(w, `{ "message": "%s" }`, err.Error())
			log.Printf("[ERROR] Encode PolicyReportResults: %s", html.EscapeString(err.Error()))
		}
	}
}

func ConfigHandler(conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		err := json.NewEncoder(w).Encode(conf)
		if err != nil {
			fmt.Fprintf(w, `{ "message": "%s" }`, err.Error())
			log.Printf("[ERROR] Encode Configuration: %s", html.EscapeString(err.Error()))
		}
	}
}
