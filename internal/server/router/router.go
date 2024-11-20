package router

import (
	"net/http"
)

type IRouter interface {
	RegisterRoutes(*http.ServeMux)
}
