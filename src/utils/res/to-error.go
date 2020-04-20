package res

import (
	"audit/src/config"
	"audit/src/utils"
	"net/http"
)

// ToError sends error
func ToError(res http.ResponseWriter, status int, err error, msg ...string) error {
	cfg := config.GetCurrentConfig()
	var appErr *utils.AppError

	switch e := err.(type) {
	case *utils.AppError:
		appErr = e
	default:
		code := "APP_ERROR"
		msgCount := len(msg)

		if msgCount > 0 {
			code = msg[0]
			msg = msg[1:]
		}

		appErr = &utils.AppError{
			Code:    code,
			Message: err.Error(),
			Err:     err,
		}
	}

	if cfg.IsDev() {
		appErr.Stack = utils.GetErrorStack(err, 2)
	}

	return ToJSON(res, status, appErr)
}
