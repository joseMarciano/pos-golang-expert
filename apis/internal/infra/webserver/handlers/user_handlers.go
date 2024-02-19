package handlers

import (
	"encoding/json"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/dto"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/user"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/infra/database"
	"net/http"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var usr dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := user.NewUser(usr.Name, usr.Email, usr.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
