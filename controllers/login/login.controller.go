package loginController

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"GO_PROJECT/config"
	"GO_PROJECT/constants"
	"GO_PROJECT/db/cassandra"
	"GO_PROJECT/db/redis"
	"GO_PROJECT/helper/jwt"
	"GO_PROJECT/middlewares"
	"GO_PROJECT/model"

	"github.com/alexedwards/argon2id"
	"github.com/go-playground/validator/v10"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Handle validation
	var loginInput model.LoginInput

	err := json.NewDecoder(r.Body).Decode(&loginInput)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(loginInput)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	loginInput.EmailId = strings.ToLower(loginInput.EmailId)

	// Check if user already registered
	user, err := cassandra.GetUser(loginInput.EmailId)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}
	if _, ok := user["user_id"]; !ok {
		middlewares.HandleResponse(w, constants.ResponseConstants.Login.UserNotExist.HTTP_STATUS, constants.ResponseConstants.Login.UserNotExist.MESSAGE, nil, err)
		return
	}

	// to get user data in it's proper format
	var userData model.User
	jsonData, err := json.Marshal(user)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}
	err = json.Unmarshal(jsonData, &userData)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}

	// check if password correct or not
	match, err := argon2id.ComparePasswordAndHash(loginInput.Password, userData.Password)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}
	if(!match){
		middlewares.HandleResponse(w, constants.ResponseConstants.Login.InvalidCredentials.HTTP_STATUS, constants.ResponseConstants.Login.InvalidCredentials.MESSAGE, nil, err)
		return
	}

	// create jwt token
	tokenData := model.JwtData{
		UserId:  userData.UserId,
		EmailId: userData.EmailId,
		Role:    userData.Role,
	}

	jsonWebToken, err := jwt.CreateToken(tokenData)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}

	// Store token in redis
	redis.SetKeyWithExpiry(constants.RedisKeys.UserAccessToken+"::"+userData.EmailId, jsonWebToken, time.Duration(config.Config.JWT_TOKEN.EXPIRY_TIME_IN_SECONDS))

	middlewares.HandleResponse(w, constants.ResponseConstants.Login.LoginSuccess.HTTP_STATUS, constants.ResponseConstants.Login.LoginSuccess.MESSAGE, map[string]interface{}{"token": jsonWebToken}, nil)
}
