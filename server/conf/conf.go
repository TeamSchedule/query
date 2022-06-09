package conf

import (
	"log"
	"os"
)

var ServerPort = os.Getenv("SERVER_PORT")

func init() {
	if ServerPort == "" {
		log.Fatal("Environment variable `SERVER_PORT` must be specified")
	}
}
