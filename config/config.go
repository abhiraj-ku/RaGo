package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL string
	APIKey  string
	Model   string
	// SystemPromptFile is the path to a text/markdown file whose
	// contents become the conversation's system message. A missing
	// file is silently treated as "no system prompt".
	SystemPromptFile string
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		BaseURL:          os.Getenv("OPENAI_BASE_URL"),
		APIKey:           os.Getenv("OPENAI_API_KEY"),
		Model:            os.Getenv("OPENAI_MODEL"),
		SystemPromptFile: os.Getenv("SYSTEM_PROMPT_FILE"),
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.openai.com/v1"
	}

	if cfg.Model == "" {
		cfg.Model = "gpt-4o-mini"
	}

	return cfg
}
