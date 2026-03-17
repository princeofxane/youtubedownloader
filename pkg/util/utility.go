package util

import "os"

func GetEnv(env string) (string, bool) {
	val := os.Getenv(env)
	if val == "" {
		return "", false
	}

	return val, true
}
