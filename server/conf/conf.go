package conf

import (
	"log"
	"os"
)

var ServerPort = os.Getenv("SERVER_PORT")
var ScheduleHosts = "http://" + os.Getenv("SCHEDULE_HOSTS")

func init() {
	if ServerPort == "" {
		log.Fatal("Environment variable `SERVER_PORT` must be specified")
	}
	if ScheduleHosts == "" {
		log.Fatal("Environment variable `SCHEDULE_HOSTS` must be specified")
	}
}
