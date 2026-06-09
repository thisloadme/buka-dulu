package config

import "fmt"

// LLMConfig holds all LLM provider configuration
type LLMConfig struct {
	// Active provider: openai, anthropic, ollama, opencode, openrouter, groq, mock, etc.
	Provider string

	// Base URL for the provider API
	BaseURL string

	// API key (not needed for Ollama/local)
	APIKey string

	// Model name
	Model string
}

// ParseLLMConfig reads LLM config from flat env vars
// Supports: openai, anthropic, ollama, opencode, openrouter, groq, together, mock
func (c *Config) ParseLLMConfig() LLMConfig {
	provider := c.LLMProvider

	// Determine base URL based on provider
	baseURL := c.LLMBaseURL
	if baseURL == "" {
		switch provider {
		case "openai":
			baseURL = "https://api.openai.com/v1"
		case "anthropic":
			baseURL = "https://api.anthropic.com/v1"
		case "openrouter":
			baseURL = "https://openrouter.ai/api/v1"
		case "groq":
			baseURL = "https://api.groq.com/openai/v1"
		case "together":
			baseURL = "https://api.together.xyz/v1"
		case "opencode":
			baseURL = "http://localhost:8080/v1" // OpenCode default
		case "ollama":
			baseURL = "http://localhost:11434"
		case "mock":
			provider = "mock"
			baseURL = ""
		default:
			// Custom provider — user must provide base URL
			fmt.Printf("WARNING: Unknown LLM provider '%s', falling back to mock\n", provider)
			provider = "mock"
		}
	}

	return LLMConfig{
		Provider: provider,
		BaseURL:  baseURL,
		APIKey:   c.LLMAPIKey,
		Model:    c.LLMModel,
	}
}
