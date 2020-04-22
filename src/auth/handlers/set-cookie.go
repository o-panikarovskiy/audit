package handlers

import (
	"audit/src/di"
	"net/http"

	"github.com/gorilla/securecookie"
)

func setAuthCookie(w http.ResponseWriter, sid string) {
	cfg := di.GetAppConfig()

	hashKey := []byte(cfg.Cookie.Hash)
	blockKey := []byte(cfg.Cookie.Block)
	s := securecookie.New(hashKey, blockKey)

	if encoded, err := s.Encode("sid", sid); err == nil {
		cookie := &http.Cookie{
			Name:  "sid",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}
