package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zHenriqueGN/EasyProduct/internal/entity"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/database"
)

func TestUserRepository_Create(t *testing.T) {
	DB := database.ConnectToMySQLDB(testDBConn)
	defer DB.Close()
	repository := NewUserRepository(DB)
	user, err := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(user)
	assert.Nil(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	DB := database.ConnectToMySQLDB(testDBConn)
	defer DB.Close()
	repository := NewUserRepository(DB)
	user, err := entity.NewUser("Kevin Rue", "kevin.rue@example.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(user)
	if err != nil {
		t.Fatal(err)
	}
	userFound, err := repository.FindByEmail("kevin.rue@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}
