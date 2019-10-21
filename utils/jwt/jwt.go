package jwt

import (
	"fmt"
	"os"
	"strings"
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

// GenerateTokenPair generate a new token pair
func GenerateTokenPair(u *models.User) (*TokenPair, error) {
	refreshClaims := jwt.StandardClaims{Id: u.ID.String(), ExpiresAt: time.Now().Add(jwtRefreshTTL).Unix()}
	accessClaims := &Claims{
		Id:         u.ID.String(),
		Privileges: strings.Fields(u.Privileges),
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

// ValidateToken ...
func ValidateToken(c buffalo.Context) (*jwt.Token, error) {
	// Get the user provided token
	tknString, err := TokenFromHeader(c)
	if err != nil {
		return &jwt.Token{}, fmt.Errorf("bad request")
	}

	// Parse and return the token
	return jwt.Parse(
		tknString,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
}

// ClaimsFromHeader validtes the token in the request and
// returns the embeded claims
func ClaimsFromHeader(c buffalo.Context) (*Claims, error) {
	var claims = new(Claims)

	// Get the user provided token
	tknString, err := TokenFromHeader(c)
	if err != nil {
		return &Claims{}, fmt.Errorf("bad request")
	}

	// Validate the token
	tkn, err := jwt.ParseWithClaims(
		tknString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return &Claims{}, fmt.Errorf("bad request")
	}

	if !tkn.Valid {
		return &Claims{}, fmt.Errorf("token invalid")
	}

	return claims, nil
}

// TokenFromHeader gets a jwt token in the form of
// Authorization: Bearer <jwt-token>
func TokenFromHeader(c buffalo.Context) (string, error) {
	tknString := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(tknString, "Bearer")

	if len(splitToken) != 2 {
		return "", fmt.Errorf("bad request")
	}

	reqToken := strings.TrimSpace(splitToken[1])
	return reqToken, nil
}
