package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"sync"
	"time"
)

var cUsers = map[string]bool{}
var CookieName = "_cookie"
var cMutes sync.Mutex

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func WriteCookie(w http.ResponseWriter) error {
	// ignoring err
	u, _ := createUUID()
	for {
		if _, ok := cUsers[u]; !ok {
			break
		}
		u, _ = createUUID()
	}

	sevenDays := 7 * 24 * 60 * 60
	c := http.Cookie{
		Name:     CookieName,
		Value:    u,
		Path:     "/",
		MaxAge:   sevenDays,
		Expires:  time.Now().Add(time.Second * time.Duration(sevenDays)),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	c.Value = base64.URLEncoding.EncodeToString([]byte(c.Value))

	if len(c.String()) > 4096 {
		return ErrValueTooLong
	}

	http.SetCookie(w, &c)
	cMutes.Lock()
	cUsers[u] = true
	cMutes.Unlock()

	return nil
}

func ReadCookie(r *http.Request, name string) error {
	cookie, err := r.Cookie(name)
	if err != nil {
		return err
	}
	value, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return ErrInvalidValue
	}
	if cUsers[string(value)] {
		return nil
	}
	return ErrInvalidValue
}
