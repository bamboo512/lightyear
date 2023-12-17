package schema

type Config struct {
	Redis    RedisConfig    `yaml:"redis"`
	Server   ServerConfig   `yaml:"server"`
	Logger   LoggerConfig   `yaml:"logger"`
	Database DatabaseConfig `yaml:"database"`
	Oss      OssConfig      `yaml:"oss"`
	Storage  StorageConfig  `yaml:"storage"`
}
