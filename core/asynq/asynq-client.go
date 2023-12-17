package asynq

import (
	"lightyear/core/global"

	"github.com/hibiken/asynq"
)

func InitClient() {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     global.Config.Redis.Address,
		Username: global.Config.Redis.Username,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})

	global.AsynqClient = client
}
