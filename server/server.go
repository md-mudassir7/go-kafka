// nolint:staticcheck,funlen,ineffassign,misspell,govet,dupl,unused,gosimple,gocritic,gocyclo,errcheck,revive,bodyclose,unconvert
package server

import (
	"context"
	"net/http"
	"sync"

	"log"

	"github.com/gorilla/mux"
	"github.com/md-mudassir7/go-kafka/pkg/consumer"
)

var (
	server *Server
	once   sync.Once
)

type Server struct {
	srv *http.Server
	r   *mux.Router
}

func init() {
	once.Do(func() {
		server = &Server{}
		server.r = mux.NewRouter()
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am healthy!!!!"))
}

func Start() {
	// initialise tracer

	server.r = mux.NewRouter()
	setAllRoutes(server.r)
	log.Println("Starting server")
	server.srv = &http.Server{Addr: ":3000", Handler: server.r}

	go func() {
		err := server.srv.ListenAndServe()
		log.Println("server.srv.ListenAndServe Failed", "error", err)
	}()

	log.Println("Kafka consumer is running...")
	consumer.ConsumeMessages()
}

func Stop() {
	server.srv.Shutdown(context.Background())
}
