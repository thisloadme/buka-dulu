package config

import (
	"bufio"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	loadDotEnv()

	cfg := &Config{
		Port:        getEnvInt("PORT", 8080),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", ""),
		JWTExpiry:   getEnvInt("JWT_EXPIRY_HOURS", 24),

		LLMProvider: getEnv("LLM_PROVIDER", "mock"),
		LLMBaseURL:  getEnv("LLM_BASE_URL", ""),
		LLMAPIKey:   getEnv("LLM_API_KEY", ""),
		LLMModel:    getEnv("LLM_MODEL", "gpt-4o-mini"),

		StorageType: getEnv("STORAGE_TYPE", "local"),
		StoragePath: getEnv("STORAGE_PATH", "./storage/evidence"),
	}

	if cfg.DatabaseURL == "" {
		slog.Error("DATABASE_URL is required — set via .env or environment variable")
		os.Exit(1)
	}
	if cfg.JWTSecret == "" {
		slog.Error("JWT_SECRET is required — set via .env or environment variable")
		os.Exit(1)
	}

	// Mask password in logs
	sanitized := cfg.DatabaseURL
	if i := strings.LastIndex(sanitized, ":"); i > 10 {
		if j := strings.Index(sanitized[i:], "@"); j > 0 {
			sanitized = sanitized[:i+1] + "***" + sanitized[i+j:]
		}
	}

	slog.Info("configuration loaded",
		"port", cfg.Port,
		"db", sanitized,
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

// loadDotEnv reads .env file if it exists and sets env vars.
// Does NOT override vars already set in the environment.
func loadDotEnv() {
	paths := []string{".env", "../.env"}
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			key, val, ok := strings.Cut(line, "=")
			if !ok {
				continue
			}
			key = strings.TrimSpace(key)
			val = strings.TrimSpace(val)
			// Don't override already-set env vars
			if os.Getenv(key) == "" {
				os.Setenv(key, val)
			}
		}
		break // found and loaded
	}
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
