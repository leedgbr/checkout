package checkout

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router http.Handler
}

func (a *App) Init() {
	r := mux.NewRouter()
	r.HandleFunc("/checkout", http.HandlerFunc(NewSimpleHandler())).
		Methods("POST").
		Headers("Content-Type", "application/json")
	a.Router = r
}
