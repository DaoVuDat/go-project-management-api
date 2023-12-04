package httpuseracc

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"project-management/domain"
)

type AccountUserHandler struct {
	AccountUserUC domain.AccountUseCase
}

func SetupAccountUserHandler(router chi.Router, accountUserHandler domain.AccountUseCase) {
	handler := &AccountUserHandler{
		AccountUserUC: accountUserHandler,
	}
	router.Route("/account/", func(r chi.Router) {
		r.Get("/test", handler.TestHandler)
		r.Post("/create-account", handler.CreateUserAccountHandler)
	})

}

/*
	Handlers
*/

func (handler *AccountUserHandler) CreateUserAccountHandler(w http.ResponseWriter, r *http.Request) {

}

func (handler *AccountUserHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"Test": "Hello",
	})
}
