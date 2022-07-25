package config

import (
	"log"
	"os"
)

type EnvironmentVariables struct {
	JWTSecret string
}

var EnvironmentVariablesData EnvironmentVariables

func LoadEnvVariables() EnvironmentVariables {
	log.Println("loading env variable")

	EnvironmentVariablesData = EnvironmentVariables{os.Getenv("JWT_SECRET")}
	return EnvironmentVariablesData
}
