package app

import (
	"net/http"
	"soccerq/internal/service"
)

type UserHandler interface {
	GetAllUser(w http.ResponseWriter, req *http.Request)
	GetUser(w http.ResponseWriter, req *http.Request)
}

type handler struct {
	Service service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &handler{Service: userService}
}

func (handler *handler) GetAllUser(w http.ResponseWriter, req *http.Request) {}

func (handler *handler) GetUser(w http.ResponseWriter, req *http.Request) {}
