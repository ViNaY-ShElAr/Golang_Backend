package signupController

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
	"github.com/gocql/gocql"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	// Handle validation
	var signupInput model.SignupInput

	err := json.NewDecoder(r.Body).Decode(&signupInput)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(signupInput)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	signupInput.EmailId = strings.ToLower(signupInput.EmailId)

	// Check if user already registered
	user, err := cassandra.GetUser(signupInput.EmailId)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}
	if _, ok := user["user_id"]; ok {
		middlewares.HandleResponse(w, constants.ResponseConstants.Signup.AlreadySignup.HTTP_STATUS, constants.ResponseConstants.Signup.AlreadySignup.MESSAGE, nil, err)
		return
	}

	// Encrpyt pass
	hashedPassword, err := argon2id.CreateHash(signupInput.Password, argon2id.DefaultParams)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}

	role := 2
	userData := model.User{
		UserId:    gocql.TimeUUID(),
		UserName:  signupInput.UserName,
		EmailId:   signupInput.EmailId,
		ContactNo: signupInput.ContactNo,
		Gender:    signupInput.Gender,
		Password:  hashedPassword,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// add User
	err = cassandra.AddUser(userData)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
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

	middlewares.HandleResponse(w, constants.ResponseConstants.Signup.SignupSuccess.HTTP_STATUS, constants.ResponseConstants.Signup.SignupSuccess.MESSAGE, map[string]interface{}{"userData": userData, "token": jsonWebToken}, nil)
}
