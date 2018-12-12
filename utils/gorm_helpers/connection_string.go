package gorm_helpers

import (
	"fmt"
	"strings"
)

const Postgres = "postgres"

func BuildConnectionString(db string, config map[string]string) (string, error) {
	switch db {
	case Postgres:
		defaultConfig := map[string]string{
			"sslmode":  "disable",
			"host":     "localhost",
			"port":     "5432",
			"user":     "postgres",
			"dbname":   "postgres",
			"password": "postgres",
		}
		return pgConnectionString(mergeDicts(defaultConfig, config)), nil
	}

	return "", fmt.Errorf("can't build connection string for database %s", db)
}

func pgConnectionString(values map[string]string) string {
	pairs := make([]string, len(values))
	i := 0
	for key, value := range values {
		pairs[i] = key + "=" + value
		i++
	}
	return strings.Join(pairs, " ")
}

func mergeDicts(dicts ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, dict := range dicts {
		for k, v := range dict {
			result[k] = v
		}
	}
	return result
}
