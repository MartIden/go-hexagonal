package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/MartIden/go-hexagonal/core/store/sqlstore"
)

type MuxServer struct {
	router *mux.Router
	logger *logrus.Logger
	Store  sqlstore.Store
}

func NewServer(store sqlstore.Store) *MuxServer {
	s := &MuxServer{
		router: mux.NewRouter(),
		logger: logrus.New(),
		Store:  store,
	}

	s.configureRouter()
	return s
}

func (srv *MuxServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}
func (srv *MuxServer) configureRouter() {

}