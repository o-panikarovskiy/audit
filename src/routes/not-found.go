package routes

import (
	"audit/src/utils"
	"net/http"
)

func notFound(res http.ResponseWriter, req *http.Request) {
	utils.ToError(res, http.StatusNotFound, &utils.AppError{
		Code:    "NOT_FOUND",
		Message: "Path not found",
	})
}
