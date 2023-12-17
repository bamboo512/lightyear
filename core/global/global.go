package global

import (
	"lightyear/core/config/schema"
	"lightyear/pkg/oss"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config      *schema.Config
	DB          *gorm.DB
	Logger      *logrus.Logger
	Redis       *redis.Client
	OssClient   *oss.OssClient
	AsynqClient *asynq.Client
)
