package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/dto"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/user"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB       database.UserInterface
	JwtAuth      *jwtauth.JWTAuth
	JwtExpiresIn int
}

type Error struct {
	Error string `json:"error"`
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		JwtAuth:      jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var userDto dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(userDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(userDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, err := h.JwtAuth.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"access_token": tokenString,
	})
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUserInput  true  "user request"
// @Success      201  {object}  user.User
// @Failure      500  {object}  Error
// @Router       /users [post]
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
