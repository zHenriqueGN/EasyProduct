package repository

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zHenriqueGN/EasyProduct/internal/entity"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/database"
)

func TruncatUsersTable(DB *sql.DB) error {
	stmt, err := DB.Prepare("TRUNCATE TABLE users")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func TestUserRepository_Create(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewUserRepository(DB)
	err := TruncatUsersTable(DB)
	if err != nil {
		t.Fatal(err)
	}
	user, err := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(user)
	assert.Nil(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewUserRepository(DB)
	err := TruncatUsersTable(DB)
	if err != nil {
		t.Fatal(err)
	}
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

func TestUserRepository_FindByEmailWhenEmailDoesNotExist(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewUserRepository(DB)
	err := TruncatUsersTable(DB)
	if err != nil {
		t.Fatal(err)
	}
	user, err := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(user)
	if err != nil {
		t.Fatal(err)
	}
	userFound, err := repository.FindByEmail("kevin.bacon@example.com")
	assert.Nil(t, userFound)
	assert.Equal(t, ErrUserNotFound, err)

}
