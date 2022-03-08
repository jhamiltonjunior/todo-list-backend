package router

import (
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"
)

func (server *Server) ListItem() {
	listItem := controller.ListItem{}

	server.HandleFunc("/api/v1/list/{id:[0-9]+}/item", listItem.CreateListItem()).Methods("GET")
	server.HandleFunc("/api/v1/list/{id:[0-9]+}/item", listItem.ShowListItem()).Methods("POST")

	server.HandleFunc(
		"/api/v1/list/{id:[0-9]+}/item/{id:[0-9]+}", listItem.UpdateListItem(),
	).Methods("PUT")

	server.HandleFunc(
		"/api/v1/list/{id:[0-9]+}/item/{id:[0-9]+}", listItem.DeleteListItem(),
	).Methods("DELETE")
}
