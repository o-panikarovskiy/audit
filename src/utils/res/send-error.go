package res

import (
	"audit/src/di"
	"audit/src/utils"
	"net/http"
)

// SendStatusError sends error to http response with http status code
func SendStatusError(res http.ResponseWriter, status int, err error, msg ...string) error {
	cfg := di.GetAppConfig()

	appErr := utils.ToAppError(err, msg...)

	if cfg.IsDev() {
		appErr.Stack = utils.GetErrorStack(err, 2)
	}

	return SendJSON(res, status, appErr)
}
