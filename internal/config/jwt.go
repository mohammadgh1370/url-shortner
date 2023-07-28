package config

import "strconv"

var (
	JWT_SECRET_KEY          = Get("JWT_SECRET_KEY", "secretkey")
	JWT_ACCESS_EXP_TIME, _  = strconv.Atoi(Get("JWT_ACCESS_EXP_TIME_IN_MINUTE", "15"))
	JWT_REFRESH_EXP_TIME, _ = strconv.Atoi(Get("JWT_REFRESH_EXP_TIME_IN_MINUTE", "86400"))
)
