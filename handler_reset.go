package main

import "net/http"


func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r * http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment."))
		return
	}
	cfg.fileserverHits.Store(0)

	if err := cfg.db.Reset(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to reset database: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and database reset to initial state."))
}

