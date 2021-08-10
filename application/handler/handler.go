package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	c "rpolnx.com.br/golang-hex/application/config"
	"rpolnx.com.br/golang-hex/application/controller"
	"rpolnx.com.br/golang-hex/application/routes"
	s "rpolnx.com.br/golang-hex/domain/service/impl"
	mRepo "rpolnx.com.br/golang-hex/infrastructure/adapter/mongo"
	rRepo "rpolnx.com.br/golang-hex/infrastructure/adapter/redis"
)

func LoadServer(config *c.Configuration) (m http.Handler, err error) {
	mongo, err := mRepo.InitializeRepo(config.Mongo)

	if err != nil {
		log.Fatal("Error initializing mongo", err)
		return nil, err
	}

	redis, err := rRepo.InitializeRepo(config.Redis)

	if err != nil {
		log.Fatal("Error initializing redis", err)
		return nil, err
	}

	service := s.NewUserService(mongo, redis)
	controller := controller.NewUserController(service)

	router := chi.NewRouter()
	c.ApplyMiddlewares(router, config.App)

	router.Route("/users", routes.AppendUserRoutes(controller))

	return router, err
}
