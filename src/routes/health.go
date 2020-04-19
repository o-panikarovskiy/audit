package routes

import (
	"audit/src/utils"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, http.StatusOK, utils.StringMap{"ok": true})
}
