package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"http-rest-api/internal/libs/calculator"
	"net/http"
)

type APIServer struct {
	serverconfig *Config
	logger       *logrus.Logger
	router       *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		serverconfig: config,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureLogger()
	s.configureRouter()

	s.logger.Info("starting API server...")

	return http.ListenAndServe(s.serverconfig.BinAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.serverconfig.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/{id}", s.handleRequest())
}

func (s *APIServer) handleRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		input := mux.Vars(r)
		output, _, err := calculator.GetAmicableNumber(input["id"])
		if err != nil {
			output = "Error in handleRequest function"
		}
		json.NewEncoder(w).Encode(output)
	}
}
