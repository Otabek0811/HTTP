package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServerHost string
	HTTPPort   string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresDB       string
	PostgresPassword string

	DefaultOffset int
	DefaultLimit  int
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST"))
	cfg.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT"))
	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER" ))
	cfg.PostgresDB = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT"))

	return cfg

}

func getOrReturnDefaultValue(key string) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		if key=="HTTP_PORT"{
			val=":"+val
		}
		return val
	}
	

	return nil
}
