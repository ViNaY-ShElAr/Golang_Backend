package logoutController

import (
	"encoding/json"
	"net/http"
	"strings"

	"GO_PROJECT/constants"
	"GO_PROJECT/db/redis"
	"GO_PROJECT/logger"
	"GO_PROJECT/middlewares"
	"GO_PROJECT/model"

	"github.com/go-playground/validator/v10"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	// Handle validation
	var logoutInput model.LogoutInput

	err := json.NewDecoder(r.Body).Decode(&logoutInput)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(logoutInput)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	// Log Request
	logger.Log.Info("requestedBy: ", r.Context().Value("accessToken").(model.JwtData))

	logoutInput.EmailId = strings.ToLower(logoutInput.EmailId)

	// Remove token from redis
	err = redis.DeleteKey(constants.RedisKeys.UserAccessToken+"::"+logoutInput.EmailId)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}

	middlewares.HandleResponse(w, constants.ResponseConstants.Logout.LogoutSuccess.HTTP_STATUS, constants.ResponseConstants.Logout.LogoutSuccess.MESSAGE, nil, nil)
}
