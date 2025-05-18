package dbconfig

import (
	auth_db "api-gateway/config/db_config/auth_db"
)

func InitDbConfig() {
	auth_db.InitAuthDbConfig()
}
