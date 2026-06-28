package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/nios-x/articles-go/internal/storage"
	"github.com/nios-x/articles-go/internal/types"
	"github.com/nios-x/articles-go/internal/util/response"
)

func New(storage storage.Storage) http.HandlerFunc {
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
		lastId, err := storage.CreateUser(
			user.Name,
			user.Email,
			user.Age,
		)

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})

	}
}

func GetByID(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid user id")))
			return
		}

		user, err := storage.GetUserByID(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("user not found")))
				return
			}
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, user)
	}
}
