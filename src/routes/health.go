package routes

import (
	"audit/src/utils"
	"audit/src/utils/res"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	res.ToJSON(w, http.StatusOK, utils.StringMap{"ok": true})
}
