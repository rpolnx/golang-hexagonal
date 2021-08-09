package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	c "rpolnx.com.br/mongo-hex/application/config"
	"rpolnx.com.br/mongo-hex/application/controller"
	"rpolnx.com.br/mongo-hex/application/routes"
	s "rpolnx.com.br/mongo-hex/domain/service/impl"
	r "rpolnx.com.br/mongo-hex/infrastructure/repository"
)

func LoadServer(config *c.Configuration) (m http.Handler, err error) {
	repo, err := r.InitializeRepo(config.Mongo)

	if err != nil {
		log.Fatal("Error initializing mongo", err)
		return nil, err
	}

	service := s.NewUserService(repo)
	controller := controller.NewUserController(service)

	router := chi.NewRouter()
	applyMiddlewares(router, config.App)

	router.Route("/users", routes.AppendUserRoutes(controller))

	return router, err
}

func applyMiddlewares(router *chi.Mux, config c.App) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))
}
