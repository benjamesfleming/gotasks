package cookies

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/securecookie"
)

var hashKey = []byte(os.Getenv("SESSION_SECRET"))
var s = securecookie.New(hashKey, hashKey)

func Set(c buffalo.Context, name string, value string) {
	encoded, err := s.Encode(name, value)
	if err == nil {
		cookie := &http.Cookie{
			Name:     name,
			Value:    encoded,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Response(), cookie)
		fmt.Fprintln(c.Response(), encoded)
	}
}

func Get(c buffalo.Context, name string) (string, error) {
	if cookie, err := c.Request().Cookie(name); err == nil {
		var value string
		if err = s.Decode(name, cookie.Value, &value); err == nil {
			fmt.Fprintln(c.Response(), value)
			return value, nil
		}
	}
	return "", nil
}
