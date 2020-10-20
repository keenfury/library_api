package config

import (
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/client9/reopen"
	"github.com/kardianos/osext"
)

var (
	AppName      = getEnvOrDefault("APP_NAME", "your_app")
	AppVersion   = getEnvOrDefault("APP_VERSION", "1.0.0")
	RestPort     = getEnvOrDefault("APP_PORT", "12572")
	PidPath      = getEnvOrDefault("PID_PATH", fmt.Sprintf("/tmp/%s.pid", AppName))
	PostgresUser = getEnvOrDefault("PG_USER", "")
	PostgresPass = getEnvOrDefault("PG_PASS", "")
	PostgresDB   = getEnvOrDefault("PG_DB", "")
	PostgresHost = getEnvOrDefault("PG_HOST", "localhost")
	LogPath      = getEnvOrDefault("LOG_PATH", fmt.Sprintf("/tmp/%s.out", AppName))
	LogOutput    *reopen.FileWriter
	ExecDir      = ""
	Env          = getEnvOrDefault("ENV", "dev")
)

func init() {
	ExecDir, _ = osext.ExecutableFolder()

	InitializeLogging()
}

func getEnvOrDefault(envVar string, defEnvVar string) (newEnvVar string) {
	if newEnvVar = os.Getenv(envVar); len(newEnvVar) == 0 {
		return defEnvVar
	} else {
		return newEnvVar
	}
}

func InitializeLogging() {
	var err error
	if LogOutput == nil {
		LogOutput, err = reopen.NewFileWriter(LogPath)
		if err != nil {
			log.Fatalf("Log output file was not set: %s", err)
		}

		// set up log format
		logFormat := &log.JSONFormatter{}
		logFormat.TimestampFormat = time.RFC3339Nano

		log.SetOutput(LogOutput)
		log.SetFormatter(logFormat)
	}
}
