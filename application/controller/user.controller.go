package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	ce "rpolnx.com.br/golang-hex/application/error"
	u "rpolnx.com.br/golang-hex/domain/model/user"
	"rpolnx.com.br/golang-hex/domain/service"
)

type UserController interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type controller struct {
	c service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &controller{
		userService,
	}
}

func (c *controller) Get(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	fmt.Println(name)

	if name == "" {
		// buildResponse(http.StatusText(http.StatusBadRequest))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := c.c.Get(name)
	if err != nil {
		if errors.Cause(err) == ce.ErrNotFound {
			fmt.Println(err)

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *controller) Post(w http.ResponseWriter, r *http.Request) {
	var user u.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.c.Post(&user)

	if err != nil {
		if errors.Cause(err) == ce.ErrInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	if name == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := c.c.Delete(name)
	if err != nil {
		if errors.Cause(err) == ce.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
}

func buildResponse(i interface{}) []byte {
	res, err := json.Marshal(i)

	if err != nil {
		log.Fatal(errors.Wrap(err, "user.controller.buildResponse"))
	}

	return res
}
