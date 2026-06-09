package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Message represents a chat message for any LLM provider
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest is the standard OpenAI-compatible request format
type ChatRequest struct {
	Model       string           `json:"model"`
	Messages    []Message        `json:"messages"`
	Temperature float64          `json:"temperature,omitempty"`
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

// ChatResponse is the standard OpenAI-compatible response format
type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// LLMProvider defines the interface for any LLM provider
type LLMProvider interface {
	Chat(req ChatRequest) (string, error)
	Name() string
}

// OpenAICompatibleProvider works with any OpenAI-compatible API:
// OpenAI, Anthropic (via OpenRouter/API), OpenRouter, Groq, Together, etc.
type OpenAICompatibleProvider struct {
	name    string
	baseURL string
	apiKey  string
	model   string
	client  *http.Client
}

func NewOpenAIProvider(name, baseURL, apiKey, model string) *OpenAICompatibleProvider {
	return &OpenAICompatibleProvider{
		name:    name,
		baseURL: baseURL,
		apiKey:  apiKey,
		model:   model,
		client:  &http.Client{Timeout: 60 * time.Second},
	}
}

func (p *OpenAICompatibleProvider) Name() string { return p.name }

func (p *OpenAICompatibleProvider) Chat(req ChatRequest) (string, error) {
	if req.Model == "" {
		req.Model = p.model
	}
	if req.Temperature == 0 {
		req.Temperature = 0.3
	}

	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	endpoint := p.baseURL + "/chat/completions"
	httpReq, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("API call: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}

	if chatResp.Error != nil {
		return "", fmt.Errorf("%s API error: %s", p.name, chatResp.Error.Message)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("%s returned no choices", p.name)
	}

	return chatResp.Choices[0].Message.Content, nil
}

// OllamaProvider uses Ollama's native API (not OpenAI-compatible mode)
type OllamaProvider struct {
	name    string
	baseURL string
	model   string
	client  *http.Client
}

func NewOllamaProvider(name, baseURL, model string) *OllamaProvider {
	return &OllamaProvider{
		name:    name,
		baseURL: baseURL,
		model:   model,
		client:  &http.Client{Timeout: 120 * time.Second},
	}
}

func (p *OllamaProvider) Name() string { return p.name }

type ollamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ollamaResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
	Done bool `json:"done"`
}

func (p *OllamaProvider) Chat(req ChatRequest) (string, error) {
	model := req.Model
	if model == "" {
		model = p.model
	}

	ollamaReq := ollamaRequest{
		Model:    model,
		Messages: req.Messages,
		Stream:   false,
	}

	body, err := json.Marshal(ollamaReq)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	endpoint := p.baseURL + "/api/chat"
	httpReq, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("Ollama API call: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	var ollamaResp ollamaResponse
	if err := json.Unmarshal(respBody, &ollamaResp); err != nil {
		return "", fmt.Errorf("parse Ollama response: %w", err)
	}

	if ollamaResp.Message.Content == "" {
		return "", fmt.Errorf("Ollama returned empty response")
	}

	return ollamaResp.Message.Content, nil
}

// MockProvider for development without any LLM
type MockProvider struct{}

func NewMockProvider() *MockProvider { return &MockProvider{} }
func (p *MockProvider) Name() string { return "mock" }

func (p *MockProvider) Chat(req ChatRequest) (string, error) {
	mock := map[string]interface{}{
		"one_line_concept":   "Bisnis F&B siap saji dengan fokus pada menu spesifik untuk target pelanggan tertentu.",
		"target_customer":    "Karyawan kantoran usia 25-40 tahun dengan budget makan siang Rp 15.000-20.000, tinggal di area perkantoran.",
		"value_proposition":  "Menu homemade dengan harga terjangkau, porsi pas, dan pengiriman cepat ke area kantor.",
		"key_assumptions":    []string{"Kantoran membutuhkan makan siang cepat", "Harga 15rb sesuai budget mayoritas", "Lokasi strategis dekat perkantoran", "Produk bisa diproduksi dari rumah"},
		"early_risks":        []string{"Persaingan dengan penjual nasi bungkus existing", "Keterbatasan kapasitas produksi rumahan"},
	}

	out, _ := json.Marshal(mock)
	return string(out), nil
}
