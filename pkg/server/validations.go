package server

import (
	"fmt"
	"strings"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
)

func validateServerConfig(serverConfig config.PostgresConfig) error {
	var errors []string
	if serverConfig.DbName == "" {
		errors = append(errors, fmt.Sprintln("DbName"))
	} else if serverConfig.Host == "" {
		errors = append(errors, fmt.Sprintln("DbHost"))
	} else if serverConfig.Port == "" {
		errors = append(errors, fmt.Sprintln("DbPort"))
	} else if serverConfig.DbUser == "" {
		errors = append(errors, fmt.Sprintln("DbUserName"))
	} else if serverConfig.DbPassword == "" {
		errors = append(errors, fmt.Sprintln("DbPassword"))
	}
	if len(errors) > 0 {
		return fmt.Errorf("%s are missing", strings.Join(errors, ","))
	}
	return nil
}
