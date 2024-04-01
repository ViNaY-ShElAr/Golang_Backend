package homeController

import (
	"encoding/json"
	"net/http"

	"GO_PROJECT/constants"
	"GO_PROJECT/kafka"
	"GO_PROJECT/logger"
	"GO_PROJECT/middlewares"
	"GO_PROJECT/model"
	"GO_PROJECT/utils"

	"github.com/go-playground/validator/v10"
)

func KafkaProducerHandler(w http.ResponseWriter, r *http.Request) {

	// Handle validation
	var messageData model.MessageData

	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(messageData)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.BAD_REQUEST.HTTP_STATUS, constants.ResponseConstants.General.BAD_REQUEST.MESSAGE, nil, err)
		return
	}

	// Log Request
	logger.Log.Info("requestedBy: ", r.Context().Value("accessToken").(model.JwtData))

	messageData.SenderId = r.Context().Value("accessToken").(model.JwtData).UserId.String()

	err = utils.StartKafkaProducer(kafka.KafkaWriter, messageData)
	if err != nil {
		middlewares.HandleResponse(w, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.HTTP_STATUS, constants.ResponseConstants.General.INTERNAL_SERVER_ERROR.MESSAGE, nil, err)
		return
	}

	middlewares.HandleResponse(w, constants.ResponseConstants.KafkaProducer.Success.HTTP_STATUS, constants.ResponseConstants.KafkaProducer.Success.MESSAGE, nil, nil)

}
