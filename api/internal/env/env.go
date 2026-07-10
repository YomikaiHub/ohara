package env

import (
	"log"
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

func MustGet(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("environment variable %s is required but not set", key)
	}
	return val
}

func MustGetInt(key string) int {
	valStr, ok := os.LookupEnv(key)
	if !ok || valStr == "" {
		log.Fatalf("environment variable %s is required but not set", key)
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Fatalf("environment variable %s must be a valid integer", key)
	}

	return val
}
