package config

import (
	"os"
)

func GetOpenAIToken() string {
	return os.Getenv("OPENAI_API_TOKEN")
}
