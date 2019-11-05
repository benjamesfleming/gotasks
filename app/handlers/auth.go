package handlers

import (
	"git.benfleming.nz/benfleming/gotasks/database/models"
	"github.com/go-pkgz/auth/token"
	"github.com/gobuffalo/nulls"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// AuthRegisterHandler handles the requests to registor a new user
// POST /auth/register
func AuthRegisterHandler(e echo.Context) error {
	db := e.Get("Database").(*gorm.DB)
	tkn := token.MustGetUserInfo(e.Request())

	user := new(models.User)
	if err := e.Bind(user); err != nil {
		return err
	}

	user.IsAdmin = false
	user.ProviderID = nulls.NewString(tkn.ID)
	user.Avatar = tkn.Picture

	if err := user.Validate(); err != nil {
		return e.JSON(400, err)
	}

	if err := db.Create(&user).Error; err != nil {
		return e.JSON(400, err)
	}

	return e.JSON(200, user)
}

// AuthMeHandler handles the request to get users own data
// GET /auth/me
func AuthMeHandler(e echo.Context) error {
	return e.JSON(200, e.Get("User").(*models.User))
}
