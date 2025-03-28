package handler

import (
	"github.com/Time-Capsule/Auth-Service/api/kafka"
	"github.com/Time-Capsule/Auth-Service/genproto/auth"
	"github.com/Time-Capsule/Auth-Service/genproto/user"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	Auth  auth.AuthServiceClient
	User  user.UserServiceClient
	Redis *redis.Client
	Kafka kafka.KafkaProducer
}

func NewHandler(auth auth.AuthServiceClient, user user.UserServiceClient, redis *redis.Client, kafka kafka.KafkaProducer) *Handler {
	return &Handler{
		Auth:  auth,
		User:  user,
		Redis: redis,
		Kafka: kafka,
	}
}
