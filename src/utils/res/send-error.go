package res

import (
	"audit/src/di"
	"audit/src/utils"
	"net/http"
)

var errorCodes = map[int]string{
	401: "AUTH_ERROR",
	403: "AUTH_ERROR",
	400: "INVALID_REQUEST_MODEL",
	500: "APP_ERROR",
}

// SendStatusError sends error to http response with http status code
func SendStatusError(res http.ResponseWriter, status int, err error, msg ...string) error {
	cfg := di.GetAppConfig()

	appErr := utils.ToAppError(err, msg...)

	if code, ok := errorCodes[status]; ok && appErr.Code == "" {
		appErr.Code = code
	}

	if cfg.IsDev() {
		appErr.Stack = utils.GetErrorStack(err, 2)
	}

	return SendJSON(res, status, appErr)
}
