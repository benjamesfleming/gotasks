package models_test

import (
	"github.com/benjamesfleming/gotasks/models"
)

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	u := &models.User{
		Email:    "test@example.com",
		UserName: "testuser",
		Password: "password123",
	}

	ms.NoError(err)
	ms.Equal(0, count)
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.Empty(verrs)
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	u := &models.User{
		UserName: "testuser",
		Password: "password123",
	}

	ms.NoError(err)
	ms.Equal(0, count)
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.NotEmpty(verrs)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	count, err := ms.DB.Count("users")
	u := &models.User{
		Email:    "test@example.com",
		UserName: "testuser",
		Password: "password123",
	}

	ms.NoError(err)
	ms.Equal(0, count)
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.Empty(verrs)
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &models.User{
		Email:    "test@example.com",
		UserName: "testuser",
		Password: "password123",
	}

	verrs, err = u.Create(ms.DB)
	ms.NoError(err)
	ms.NotEmpty(verrs)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}
