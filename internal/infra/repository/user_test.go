package repository

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zHenriqueGN/EasyProduct/internal/entity"
)

func createTestUsersTable(t *testing.T, DB *sql.DB) {
	_, err := DB.Exec(`
		CREATE TABLE users (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)`,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserRepository_Create(t *testing.T) {
	DB := NewSQLiteDB(t)
	defer DB.Close()
	createTestUsersTable(t, DB)
	repository := NewUserRepository(DB)
	user, err := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(user)
	assert.Nil(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	DB := NewSQLiteDB(t)
	defer DB.Close()
	createTestUsersTable(t, DB)
	repository := NewUserRepository(DB)
	user, err := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(user)
	if err != nil {
		t.Fatal(err)
	}
	userFound, err := repository.FindByEmail("john.doe@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}
