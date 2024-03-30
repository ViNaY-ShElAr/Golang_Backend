package model

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/golang-jwt/jwt"
)

// configurations
type Configurations struct {
	APP struct {
		PREFIX string `json:"prefix"`
	}
	JWT_TOKEN struct {
		EXPIRY_TIME            int    `json:"expiry_time"`
		EXPIRY_TIME_IN_SECONDS int    `json:"expiry_time_in_seconds"`
		TOKEN_SECRET           string `json:"token_secret"`
		TOKEN_ISSUER           string `json:"token_issuer"`
	}
	KAFKA struct {
		CONSUMER_GROUP string `json:"consumer_group"`
		CONSUMER_TOPIC string `json:"consumer_topic"`
	}
}


// secrets
type CassandraCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Keyspace string `json:"keyspace"`
}

type RedisCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}


// validate inputs
type SignupInput struct {
	UserName  string `json:"userName" validate:"required"`
	EmailId   string `json:"emailId" validate:"required,email"`
	ContactNo string `json:"contactNo,omitempty"`
	Gender    string `json:"gender" validate:"required,oneof=male female trans"`
	Password  string `json:"password" validate:"required"`
}

type LoginInput struct {
	EmailId  string `json:"emailId" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LogoutInput struct {
	EmailId string `json:"emailId" validate:"required,email"`
}


// database models
type User struct {
	UserId    gocql.UUID `json:"user_id"`
	UserName  string     `json:"user_name"`
	EmailId   string     `json:"email_id"`
	ContactNo string     `json:"contact_no"`
	Gender    string     `json:"gender"`
	Password  string     `json:"password"`
	Role      int        `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}


// general models
type ApiResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      string      `json:"error"`
}

type JwtData struct {
	UserId  gocql.UUID `json:"user_id"`
	EmailId string     `json:"email_id"`
	Role    int        `json:"role"`
	jwt.StandardClaims
}

type MessageData struct {
	SenderId string `json:"sender_id"`
	To string `json:"to"`
	Body string `json:"body"`
}