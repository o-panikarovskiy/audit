package handlers

import (
	"audit/src/di"
	"net/http"

	"github.com/gorilla/securecookie"
)

func setAuthCookie(w http.ResponseWriter, sid string) {
	cfg := di.GetAppConfig()

	name := cfg.Cookie.Name
	hashKey := []byte(cfg.Cookie.Hash)
	blockKey := []byte(cfg.Cookie.Block)

	s := securecookie.New(hashKey, blockKey)
	s.MaxAge(cfg.SessionAgeSec)

	if encoded, err := s.Encode(name, sid); err == nil {
		cookie := &http.Cookie{
			Name:  name,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}
