package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	ce "rpolnx.com.br/golang-hex/application/error"
	u "rpolnx.com.br/golang-hex/domain/model/user"
	"rpolnx.com.br/golang-hex/domain/service"
)

type UserController interface {
	GetAll(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type controller struct {
	s service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &controller{
		userService,
	}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.s.GetAll()

	if err != nil {
		if errors.Cause(err) == ce.ErrNotFound {
			log.Println(err)

			render.Render(w, r, ce.ErrResponseNotFound)
			return
		}
		render.Render(w, r, ce.ErrRender(err))
		return
	}

	body, err := json.Marshal(users)

	if err != nil {
		render.Render(w, r, ce.ErrRender(err))
		return
	}

	setupResponse(w, body, 200)
}

func (c *controller) Get(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	if name == "" {
		render.Render(w, r, ce.ErrInvalidRequest(ce.ErrInvalid))
		return
	}

	user, err := c.s.Get(name)
	if err != nil {
		if errors.Cause(err) == ce.ErrNotFound {
			log.Println(err)

			render.Render(w, r, ce.ErrResponseNotFound)
			return
		}
		render.Render(w, r, ce.ErrRender(err))
		return
	}

	body, err := json.Marshal(user)

	if err != nil {
		render.Render(w, r, ce.ErrRender(err))
		return
	}

	setupResponse(w, body, 200)
}

func (h *controller) Post(w http.ResponseWriter, r *http.Request) {
	var user u.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		render.Render(w, r, ce.ErrRender(err))
		return
	}

	err = h.s.Post(&user)

	if err != nil {
		if errors.Cause(err) == ce.ErrInvalid {
			render.Render(w, r, ce.ErrInvalidRequest(ce.ErrInvalid))
			return
		}
		render.Render(w, r, ce.ErrRender(err))
		return
	}
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	if name == "" {
		render.Render(w, r, ce.ErrInvalidRequest(ce.ErrInvalid))
		return
	}

	err := c.s.Delete(name)
	if err != nil {
		if errors.Cause(err) == ce.ErrNotFound {
			render.Render(w, r, ce.ErrResponseNotFound)
			return
		}
		render.Render(w, r, ce.ErrRender(err))
		return
	}
	render.Status(r, 204)
}

func setupResponse(w http.ResponseWriter, body []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
