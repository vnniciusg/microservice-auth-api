package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vnniciusg/microservice-auth-api/config"
)

type PingResponse struct {
	Message string `json:"message"`
}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response := PingResponse{Message: "pong"}

		jsonData, err := json.Marshal(response)

		if err != nil {
			restErr := config.NewInternalServerError("Something goes wrong")
			http.Error(w, restErr.Err, restErr.Code)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}).Methods("GET")

	return apiRouter
}
