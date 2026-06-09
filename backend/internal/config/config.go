package config

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port        int
	DatabaseURL string
	JWTSecret   string
	JWTExpiry   int

	LLMProvider string
	LLMBaseURL  string
	LLMAPIKey   string
	LLMModel    string

	StorageType string
	StoragePath string
}

func Load() *Config {
	cfg := &Config{
		Port:        getEnvInt("PORT", 8080),
		DatabaseURL: getEnv("DATABASE_URL", "sqlite://./data/bukadulu.db"),
		JWTSecret:   getEnv("JWT_SECRET", "dev-secret-change-in-prod"),
		JWTExpiry:   getEnvInt("JWT_EXPIRY_HOURS", 24),

		LLMProvider: getEnv("LLM_PROVIDER", "mock"),
		LLMBaseURL:  getEnv("LLM_BASE_URL", ""),
		LLMAPIKey:   getEnv("LLM_API_KEY", ""),
		LLMModel:    getEnv("LLM_MODEL", "gpt-4o-mini"),

		StorageType: getEnv("STORAGE_TYPE", "local"),
		StoragePath: getEnv("STORAGE_PATH", "./storage/evidence"),
	}

	slog.Info("configuration loaded",
		"port", cfg.Port,
		"db", cfg.DatabaseURL,
		"llm_provider", cfg.LLMProvider,
		"llm_model", cfg.LLMModel,
	)

	llmCfg := cfg.GetLLMConfig()
	slog.Info("llm provider configured",
		"provider", llmCfg.Provider,
		"base_url", llmCfg.BaseURL,
		"model", llmCfg.Model,
		"has_api_key", llmCfg.APIKey != "",
	)

	return cfg
}

func (c *Config) GetLLMConfig() LLMConfig {
	return c.ParseLLMConfig()
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func (c *Config) CheckLLMHealth() bool {
	llmCfg := c.ParseLLMConfig()
	if llmCfg.Provider == "mock" || llmCfg.Provider == "ollama" {
		return true
	}
	if llmCfg.BaseURL == "" {
		return false
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(llmCfg.BaseURL + "/models")
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
