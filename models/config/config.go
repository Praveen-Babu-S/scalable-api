package config

type PostgresConfig struct {
	Host       string
	Port       string
	DbName     string
	DbUser     string
	DbPassword string
}

type ServerConfig struct {
	ServerPort     string
	PostgresConfig PostgresConfig
}
