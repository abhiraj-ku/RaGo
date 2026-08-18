# RaGo - Rag in go

A Go-based learning project for building a modern Retrieval-Augmented Generation (RAG) application, grounded in clean architecture, provider abstraction, and production-minded AI engineering patterns.


## Why this project exists
AS i wanted to understand and actually see how AI systems are buitl and what exacly is RAG and whole buzz around it:

- environment-based configuration
- provider abstraction
- message history management
- prompt orchestration
- retrieval and grounding
- clean separation of concerns


## Current project status
 The code currently includes:

- configuration loading from environment variables
- OpenAI-compatible LLM client setup
- app-level wiring for a future chat / REPL flow
- package separation that will scale into retrieval and embedding layers

- chat layer is in next commits 
## Architecture overview

The project follows a simple but solid engineering pattern:

- `cmd/rag` contains the process entry point
- `app` orchestrates the runtime application flow
- `config` loads runtime configuration from environment variables
- `llm` wraps the LLM provider SDK and exposes a clean app-level interface

This keeps the system decoupled from one vendor and allows it to evolve into a broader RAG stack without forcing large rewrites.

### High-level flow

```text
main
  -> app.Run
      -> config.Load
      -> llm.New
      -> chat.RunREPL (planned / future layer)
```

This makes the app composition simple and keeps the entrypoint small and readable.

## Project structure

```text
.
├── app/
│   └── app.go                 # app orchestration and composition root
├── cmd/
│   └── rag/
│       └── main.go            # CLI entry point
├── config/
│   └── config.go              # env-based configuration
├── llm/
│   └── client.go              # OpenAI-compatible LLM client wrapper
├── go.mod
├── go.sum
└── README.md
```

## Core concepts in this project

### 1. Provider abstraction

The LLM client is designed to work with OpenAI-compatible backends, including OpenAI itself and local or self-hosted alternatives like Ollama, LM Studio, or similar services.

This is important because it keeps the app flexible and avoids locking business logic to a single vendor.

### 2. Stateless model interaction

The model does not remember prior conversation unless the app re-sends the relevant messages. This is why message history must be managed by the application layer.

### 3. Configuration-driven runtime behavior

The project intentionally reads runtime settings from environment variables so it can run cleanly across local development, containers, and deployment environments without hardcoding secrets or model choices.

### 4. Clean boundaries

By isolating config and provider logic from app startup, the project prepares for future features such as:

- chat history management
- retrieval from a vector database
- embedding generation
- RAG prompt assembly
- document ingestion
- web UI or API layer

## Configuration

The app loads settings from environment variables and optionally a `.env` file.

### Supported variables

- `OPENAI_BASE_URL` — defaults to `https://api.openai.com/v1`
- `OPENAI_API_KEY` — optional, used when provided
- `OPENAI_MODEL` — defaults to `gpt-4o-mini`
- `SYSTEM_PROMPT_FILE` — optional path to a system prompt file

### Example `.env`

```env
OPENAI_BASE_URL=https://api.openai.com/v1
OPENAI_API_KEY=your_api_key_here
OPENAI_MODEL=gpt-4o-mini
SYSTEM_PROMPT_FILE=./prompts/system.md
```

## Getting started

### Prerequisites

- Go 1.22+
- An OpenAI-compatible API key or a local provider endpoint

### Install dependencies

```bash
go mod tidy
```

### Run the app

```bash
go run ./cmd/rag
```



