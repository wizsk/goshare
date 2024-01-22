package cookie

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var cUsers = map[string]bool{}
var CookieName = "__cokief_gf"
var cMutes sync.Mutex

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func CreateUUID() (string, error) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		return "", err
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:]), nil
}

func WriteCookie(w http.ResponseWriter) error {
	// ignoring err
	u, err := CreateUUID()
	if err != nil {
		return err
	}

	for {
		if _, ok := cUsers[u]; !ok {
			break
		}
		u, err = CreateUUID()
		if err != nil {
			return err
		}
	}

	deadline := 14 * 24 * 60 * 60 // in seconds
	c := http.Cookie{
		Name:     CookieName,
		Value:    u,
		Path:     "/",
		MaxAge:   deadline,
		Expires:  time.Now().Add(time.Second * time.Duration(deadline)),
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
