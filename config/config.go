package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type UniDocConfig struct {
	Key string
}

func GetUniDocCred() (*UniDocConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	key := os.Getenv("UNIDOC_LICENSE_API_KEY")
	if len(key) == 0 {
		return nil, errors.New("uni doc key not found")
	}

	return &UniDocConfig{Key: key}, nil
}
