package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/benjamesfleming/gotasks/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
)

var jwtKey = []byte(os.Getenv("SESSION_SECRET"))
var jwtMethod = jwt.SigningMethodHS256

var jwtAccessTTL = 5 * time.Minute
var jwtRefreshTTL = 24 * time.Hour * 7

// TokenPair ...
// A single access & refresh token pair
type TokenPair struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// GenerateTokens generate a new token pair
func GenerateTokens(u *models.User) (*TokenPair, error) {
	refreshClaims := jwt.StandardClaims{ExpiresAt: time.Now().Add(jwtRefreshTTL).Unix()}
	accessClaims := &Claims{
		ID:    u.ID,
		Roles: u.Roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtAccessTTL).Unix(),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwtMethod, refreshClaims).SignedString(jwtKey)
	accessToken, err := jwt.NewWithClaims(jwtMethod, accessClaims).SignedString(jwtKey)

	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens")
	}

	return &TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

// ValidateToken validtes the token in the request and
// returns the embed claims
func ValidateToken(c buffalo.Context) (*Claims, error) {
	claims := &Claims{}

	tnkString := c.Request().Header.Get("Authentication")
	tkn, err := jwt.ParseWithClaims(tnkString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("bad request")
	}

	if !tkn.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return claims, nil
}
