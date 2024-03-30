package auth

import (
	"GO_PROJECT/constants"
	"GO_PROJECT/db/redis"
	"GO_PROJECT/helper/jwt"
	"GO_PROJECT/middlewares"
	"GO_PROJECT/model"
	"context"
	"encoding/json"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-access-token")
		if token == "" {
			middlewares.HandleResponse(w, constants.ResponseConstants.General.UNAUTHORIZED.HTTP_STATUS, constants.ResponseConstants.General.UNAUTHORIZED.MESSAGE, nil, nil)
			return
		}

		decoded, err := jwt.VerifyToken(token)
		if err != nil {
			middlewares.HandleResponse(w, constants.ResponseConstants.General.UNAUTHORIZED.HTTP_STATUS, constants.ResponseConstants.General.UNAUTHORIZED.MESSAGE, nil, err)
			return
		}
		
		// to get user data in it's proper format
		var jwtData model.JwtData
		jsonData, err := json.Marshal(decoded.Claims)
		if err != nil {
			middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
			return
		}
		err = json.Unmarshal(jsonData, &jwtData)
		if err != nil {
			middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
			return
		}
		
		// Check if user has active redis session or not
		_, err = redis.GetKey(constants.RedisKeys.UserAccessToken+"::"+jwtData.EmailId)
		if err != nil {
			middlewares.HandleResponse(w, constants.ResponseConstants.General.UNAUTHORIZED.HTTP_STATUS, constants.ResponseConstants.General.UNAUTHORIZED.MESSAGE, nil, err)
			return
		}

		// Create a context with the access token and add it to the request
		ctx := context.WithValue(r.Context(), "accessToken", jwtData)

		// Call the next handler with the modified request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
