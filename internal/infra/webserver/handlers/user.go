package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/zHenriqueGN/EasyProduct/internal/dto"
	"github.com/zHenriqueGN/EasyProduct/internal/entity"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/repository"
)

type UserHandler struct {
	UserRepository repository.UserInterface
}

func NewUserHandler(userRepository repository.UserInterface) *UserHandler {
	return &UserHandler{UserRepository: userRepository}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	user, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	err = h.UserRepository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var getJWTInput dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&getJWTInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	user, err := h.UserRepository.FindByEmail(getJWTInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	if !user.ValidatePassword(getJWTInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	JWT := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	JWTExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	_, token, err := JWT.Encode(map[string]any{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(JWTExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	accessToken := map[string]string{
		"AccessToken": token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
