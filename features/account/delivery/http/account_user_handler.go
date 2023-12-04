package httpuseracc

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
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
	ctx := r.Context()
	data := &domain.AccountCreateAndLoginRequest{}

	if err := render.Bind(r, data); err != nil {
		err := render.Render(w, r, domain.ErrInvalidRequest(err))
		if err != nil {
			log.Fatalf("Render Response failed: %s", err.Error())
			return
		}
		return
	}

	accountResponse, err := handler.AccountUserUC.CreateUserAccount(ctx, data.Username, data.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUsernameExists) {
			err := render.Render(w, r, domain.ErrResourceConflict(err))
			if err != nil {
				log.Fatalf("Render Response failed: %s", err.Error())
				return
			}
			return
		}
		err := render.Render(w, r, domain.ErrInternal(err))
		if err != nil {
			log.Fatalf("Render Response failed: %s", err.Error())
			return
		}
		return
	}

	render.Render(w, r, &accountResponse)
}

func (handler *AccountUserHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"Test": "Hello",
	})
}
