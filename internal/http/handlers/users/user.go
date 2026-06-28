package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nios-x/articles-go/internal/types"
	"github.com/nios-x/articles-go/internal/util/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err := validator.New().Struct(user); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "true"})

	}
}
