package auth

import (
	"net/http"
	"github.com/gorilla/securecookie"
)

const USER_COOKIE_NAME = "rywww"

var SecureCookie *securecookie.SecureCookie

func initCookie(secret string) {
	var hashKey = []byte(secret)
	var blockKey = []byte(nil)
	SecureCookie = securecookie.New(hashKey, blockKey)
}

func Me(r *http.Request) (int64, string) {
	userid, username := ReadUser(r)
	if userid <= 0 {
		return 0, "not login"
	}

	return userid, username
}

func ReadUser(r *http.Request) (int64, string) {
	if cookie, err := r.Cookie(USER_COOKIE_NAME); err == nil {
		var value CookieData
		if err = SecureCookie.Decode(USER_COOKIE_NAME, cookie.Value, &value); err == nil {
			return value.UserId, value.Username
		}
	}

	return 0, ""
}