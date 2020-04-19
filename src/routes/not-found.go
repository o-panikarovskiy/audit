package routes

import (
	"audit/src/utils"
	"net/http"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	panic(utils.NewAPPError(http.StatusNotFound, `NOT_FOUND`, `Path not found.`))
}
