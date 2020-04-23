package routes

import (
	"audit/src/utils/res"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	res.SendJSON(w, http.StatusOK, map[string]interface{}{"ok": true})
}
