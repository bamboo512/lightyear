package config

import (
	"fmt"
	"lightyear/core/config/schema"
	"lightyear/core/global"
	"os"

	"gopkg.in/yaml.v2"
)

func InitConfig() {
	const configPath = "config/config.yaml"
	config := &schema.Config{}

	yamlConfig, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("get settings.yaml failed: %s", err))
	}
	err = yaml.Unmarshal(yamlConfig, config)
	if err != nil {
		panic(fmt.Errorf("parse settings.yaml failed: %s", err))
	}

	global.Config = config

	// log.Println("init config successfully") // 其实不用显示 配置加载成功。因为如果没有加载成功，会直接提示错误并 panic
}
