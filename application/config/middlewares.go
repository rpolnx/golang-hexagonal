package config

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func ApplyMiddlewares(r *chi.Mux, config App) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(contextRequestID)
	r.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))
}

func contextRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		v := ctx.Value(middleware.RequestIDKey).(string)

		w.Header().Set("X-Request-Id", v)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
