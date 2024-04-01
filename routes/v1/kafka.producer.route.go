package routes

import (
	"github.com/gorilla/mux"

	kafkaProducerController "GO_PROJECT/controllers/kafkaProducer"
	"GO_PROJECT/middlewares/auth"
)

func RegisterKafkaProducerRoutes(r *mux.Router) {
	r.HandleFunc("/kafkaProducer", auth.Auth(kafkaProducerController.KafkaProducerHandler)).Methods("POST")
}
