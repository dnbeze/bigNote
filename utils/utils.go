package utils

import (
	"errors"
	"os"
)

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return "", errors.New("failed to load Env")
	}
	return value, nil
}
