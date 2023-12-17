package cmd

import (
	"lightyear/core/asynq"
	"lightyear/core/config"
	"lightyear/core/database"
	"lightyear/core/logger"
	"lightyear/core/redis"
)

func InitialApp() {
	config.InitConfig()
	logger.InitLogger()
	database.InitDatabase()
	database.MigrateDatabase()
	redis.InitRedisClient()

	go asynq.InitServer()
	asynq.InitClient()
}
