package app

import (
	"context"
	"rag-course/chat"
	"rag-course/config"
	"rag-course/llm"
)

// Run is the program's main loop
func Run(ctx context.Context, cfg config.Config) error {
	client := llm.New(cfg)
	return chat.RunREPL(ctx, client, chat.Options{
		SystemPromptFile: cfg.SystemPromptFile,
	})
}
