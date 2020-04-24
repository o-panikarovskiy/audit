package sockets

import (
	"audit/src/di"
	"audit/src/middlewares"
	"audit/src/utils/res"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

// HTTPUpgradeHandler connect socket handler
func HTTPUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	u := middlewares.GetContext(r).GetSessionUser()
	if u == nil {
		res.WriteJSONHeader(w, http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		res.SendStatusError(w, http.StatusBadRequest, err)
		return
	}

	startListen(createClient(conn, u.ID))
}

func readAuthCookie(r *http.Request) string {
	cfg := di.GetAppConfig()

	name := cfg.Cookie.Name
	hashKey := []byte(cfg.Cookie.Hash)
	blockKey := []byte(cfg.Cookie.Block)

	s := securecookie.New(hashKey, blockKey)

	if cookie, err := r.Cookie(name); err == nil {
		var value = ""
		if err = s.Decode(name, cookie.Value, &value); err == nil {
			return value
		}
	}

	return ""
}
