package config

import "os"

type Cfg struct {
	DbDsn     string
	DbDialect string
}

func New() Cfg {
	return Cfg{
		DbDsn:     readFromEnv("DB_DSN", "notes.db"),
		DbDialect: readFromEnv("DB_DIALECT", "sqlite3"),
	}
}

func readFromEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}
