package routes

import (
	"audit/src/utils"
	"audit/src/utils/res"
	"net/http"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	res.ToError(w, http.StatusNotFound, &utils.AppError{
		Code:    "NOT_FOUND",
		Message: "Path not found",
	})
}
