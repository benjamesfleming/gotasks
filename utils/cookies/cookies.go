package cookies

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/securecookie"
)

var hashKey = []byte(os.Getenv("COOKIE_HASH_KEY"))
var blockKey = []byte(os.Getenv("COOKIE_BLOCK_KEY"))

var s = securecookie.New(hashKey, blockKey)

func Set(c buffalo.Context, name string, value string) {
	encoded, err := s.Encode(name, value)
	fmt.Println(err)
	if err == nil {
		cookie := &http.Cookie{
			Name:     name,
			Value:    encoded,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Response(), cookie)
	}
}

func Get(c buffalo.Context, name string) (string, error) {
	if cookie, err := c.Request().Cookie(name); err == nil {
		var value string
		if err = s.Decode(name, cookie.Value, &value); err == nil {
			return value, nil
		}
	}
	return "", nil
}
