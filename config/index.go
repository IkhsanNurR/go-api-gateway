package config

import (
	app_config "api-gateway/config/app_config"
	db_config "api-gateway/config/db_config"
	redisconfig "api-gateway/config/redis_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDbConfig()
	redisconfig.InitRedis()
}
