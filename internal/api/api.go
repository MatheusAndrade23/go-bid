package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/matheusandrade23/go-bid/internal/services"
)

type Api struct{
	Router *chi.Mux
	UserService services.UserService
	Sessions *scs.SessionManager
}

