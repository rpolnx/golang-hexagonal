package routes

import (
	"github.com/go-chi/chi/v5"
	"rpolnx.com.br/golang-hex/application/controller"
)

func AppendUserRoutes(controller controller.UserController) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/", controller.GetAll)
		r.Get("/{name}", controller.Get)
		r.Post("/", controller.Post)
		r.Delete("/{name}", controller.Delete)
	}
}
