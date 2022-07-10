package config

import (
	"os"
)

type EnvironmentVariables struct {
	JWTSecret string
}

var EnvironmentVariablesData EnvironmentVariables

func LoadEnvVariables() {
	EnvironmentVariablesData = EnvironmentVariables{os.Getenv("JWT_SECRET")}
}
