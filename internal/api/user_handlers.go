package api

import (
	"errors"
	"net/http"

	"github.com/matheusandrade23/go-bid/internal/jsonutils"
	"github.com/matheusandrade23/go-bid/internal/services"
	"github.com/matheusandrade23/go-bid/internal/usecases/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request){
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return 
	}

	id, err := api.UserService.CreateUser(
		r.Context(), 
		data.UserName, 
		data.Email, 
		data.Bio, 
		data.Password,
	)

	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrPassword) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "email or username already exists",
			})
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request){
	
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request){
	
}