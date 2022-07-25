package common

import "os"

func GetEnv(key, fallback string) string {
	if value, defined := os.LookupEnv(key); defined {
		return value
	}
	return fallback
}
