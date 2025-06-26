# twitchNotify Project Overview for AI Agent

This document provides essential information for an AI agent working on the `twitchNotify` project.

## 1. Project Purpose
`twitchNotify` is a Go application designed to notify users about Twitch stream events (e.g., when a followed streamer goes live). It likely integrates with the Twitch API and uses a configuration file (e.g., `sample.toml`) for settings.

## 2. Technology Stack
- **Language:** Go (Golang)
- **Dependency Management:** Go Modules (`go.mod`, `go.sum`)
- **Configuration:** TOML format (e.g., `sample.toml`)

## 3. Key Directories and Files
- `main.go`: The main entry point of the application.
- `commands/`: Contains subcommands or specific functionalities (e.g., `ping.go`).
- `internal/config/config.go`: Handles application configuration loading and parsing.
- `sample.toml`: An example configuration file.
- `go.mod`, `go.sum`: Go module definition and dependency checksums.

## 4. Development Environment Setup
To set up the development environment:
- Ensure Go is installed (version specified in `go.mod` if any, otherwise latest stable).
- Dependencies are managed by Go Modules. They will be automatically downloaded when you build or run the project.

## 5. Building and Running
- **Build:** `go build -o twitchNotify .`
- **Run:** `./twitchNotify` (after building) or `go run main.go`

## 6. Testing
- **Run all tests:** `go test ./...`
- **Run tests for a specific package:** `go test ./commands`

## 7. Configuration
The application uses a TOML file for configuration. Refer to `sample.toml` for an example of the expected configuration structure. The `internal/config/config.go` file defines how this configuration is loaded and structured.

## 8. Coding Conventions
- Follow standard Go formatting (`go fmt`).
- Adhere to Go's idiomatic practices.
- Keep functions concise and focused.
- Use clear and descriptive variable/function names.

## 9. Common Tasks
- Adding new commands.
- Extending Twitch API integration.
- Modifying configuration options.
- Improving notification mechanisms.
