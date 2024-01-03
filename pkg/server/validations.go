package server

import (
	"fmt"
	"strings"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
)

func validateServerConfig(serverConfig config.ServerConfig) error {
	var errors []string
	if serverConfig.PostgresConfig.DbName == "" {
		errors = append(errors, fmt.Sprintln("DbName"))
	} else if serverConfig.PostgresConfig.Host == "" {
		errors = append(errors, fmt.Sprintln("DbHost"))
	} else if serverConfig.PostgresConfig.Port == "" {
		errors = append(errors, fmt.Sprintln("DbPort"))
	} else if serverConfig.PostgresConfig.DbUser == "" {
		errors = append(errors, fmt.Sprintln("DbUserName"))
	} else if serverConfig.PostgresConfig.DbPassword == "" {
		errors = append(errors, fmt.Sprintln("DbPassword"))
	} else if serverConfig.ServerPort == "" {
		errors = append(errors, fmt.Sprintln("ServerPort"))
	}
	if len(errors) > 0 {
		return fmt.Errorf("%s are missing", strings.Join(errors, ","))
	}
	return nil
}
