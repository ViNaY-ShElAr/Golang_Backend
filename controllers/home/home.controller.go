package homeController

import (
	"net/http"

	"GO_PROJECT/constants"
	"GO_PROJECT/logger"
	"GO_PROJECT/middlewares"
	"GO_PROJECT/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Log Request
	logger.Log.Info("requestedBy: ", r.Context().Value("accessToken").(model.JwtData))

	middlewares.HandleResponse(w, constants.ResponseConstants.Home.Success.HTTP_STATUS, constants.ResponseConstants.Home.Success.MESSAGE, nil, nil)

}
