package router

import (
	"net/http"
)

type IRouter interface {
	GetRoutes() *http.ServeMux
	GetPattern() string
}
