package handlers

import (
	"encoding/json"
	"net/http"

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
		return
	}
	user, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserRepository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
