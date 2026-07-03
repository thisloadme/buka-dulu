package service

import (
	"encoding/json"
	"fmt"

	"github.com/riyantobudi/bukadulu/internal/domain"
)

// LLMService orchestrates LLM calls using the configured provider
type LLMService struct {
	provider LLMProvider
}

// NewLLMService creates an LLM service with the given provider
func NewLLMService(provider LLMProvider) *LLMService {
	return &LLMService{provider: provider}
}

// StructureIdea sends raw input to LLM and returns a structured concept
func (s *LLMService) StructureIdea(rawInput string) (*domain.StructuredConcept, error) {
	systemPrompt := `Kamu adalah asisten yang membantu calon pebisnis F&B menyusun ide bisnis mereka.
Tugasmu adalah mengubah ide mentah menjadi konsep terstruktur dalam BAHASA INDONESIA.

Output harus JSON dengan format:
{
  "one_line_concept": "Konsep dalam satu kalimat",
  "target_customer": "Deskripsi target pelanggan",
  "value_proposition": "Nilai unik yang ditawarkan",
  "key_assumptions": ["asumsi 1", "asumsi 2", ...],
  "early_risks": ["risiko 1", "risiko 2", ...]
}

Buatlah output yang tajam, spesifik, dan tidak generik.
Key assumptions minimal 3, early risks minimal 2.`

	userPrompt := fmt.Sprintf(`Berikut ide bisnis F&B yang ingin divalidasi pengguna:

"%s"

Buatkan struktur konsep bisnis dalam format JSON yang sudah ditentukan.`, rawInput)

	req := ChatRequest{
		Temperature: 0.3,
		ResponseFormat: &ResponseFormat{Type: "json_object"},
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
	}

	content, err := s.provider.Chat(req)
	if err != nil {
		return nil, fmt.Errorf("LLM call failed: %w", err)
	}

	var concept domain.StructuredConcept
	if err := json.Unmarshal([]byte(content), &concept); err != nil {
		return nil, fmt.Errorf("parse LLM response: %w", err)
	}

	return &concept, nil
}

// RefineIdeaResult is the structured output of an AI refinement iteration.
type RefineIdeaResult struct {
	OneLineConcept   string   `json:"one_line_concept"`
	TargetCustomer   string   `json:"target_customer"`
	ValueProposition string   `json:"value_proposition"`
	KeyAssumptions   []string `json:"key_assumptions"`
	EarlyRisks       []string `json:"early_risks"`
	Summary          string   `json:"summary"`
}

// RefineIdea runs one AI iteration over the current idea concept, applying the
// user's instruction while preserving the idea's core identity. Conversation
// history is threaded so multi-turn refinement keeps context.
//
// The system prompt strictly bounds the assistant to the user's F&B business
// idea and forbids any other topic.
func (s *LLMService) RefineIdea(rawInput, currentConcept string, history []Message, instruction string) (*RefineIdeaResult, error) {
	systemPrompt := `Kamu adalah asisten validasi ide bisnis F&B (makanan & minuman) bernama BukaDulu.
Tujuanmu satu-satunya: membantu pengguna menyempurnakan dan mempertajam KONSEP IDE BISNIS F&B mereka.

ATURAN MUTLAK (tidak bisa dinegosiasikan):
1. HANYA membahas ide bisnis F&B pengguna. Setiap topik di luar itu (programming, politik, agama, topik pribadi, permintaan tulis kode, dll) TOLAK dengan sopan dan arahkan kembali ke ide F&B mereka.
2. Jawab SELALU dalam Bahasa Indonesia.
3. Pertahankan identitas inti ide (produk utama, niche). Jangan mengganti ide menjadi bisnis yang sama sekali berbeda.
4. Setiap respons HARUS berupa JSON valid dengan format:
{
  "one_line_concept": "Konsep satu kalimat (diperbarui jika perlu)",
  "target_customer": "Target pelanggan yang tajam",
  "value_proposition": "Nilai unik yang ditawarkan",
  "key_assumptions": ["asumsi 1", "asumsi 2", ...],
  "early_risks": ["risiko 1", "risiko 2", ...],
  "summary": "Ringkasan singkat perubahan/alasan dalam 1-2 kalimat"
}
5. Jangan menambahkan teks di luar JSON. Jangan bungkus dengan markdown code fence.`

	msgs := []Message{{Role: "system", Content: systemPrompt}}

	// Konteks ide asli (raw) + konsep saat ini sebagai pegangan
	contextMsg := fmt.Sprintf(`KONTEKS IDE ASLI PENGGUNA:
"%s"

KONSEP SAAT INI:
%s

Pegang konteks ide di atas sebagai batas pembahasan.`, rawInput, ifEmpty(currentConcept, "(belum ada konsep)"))
	msgs = append(msgs, Message{Role: "system", Content: contextMsg})

	// Riwayat percakapan sebelumnya (role user/assistant bergantian)
	msgs = append(msgs, history...)

	// Instruksi iterasi ini
	msgs = append(msgs, Message{Role: "user", Content: instruction})

	req := ChatRequest{
		Temperature:    0.4,
		ResponseFormat: &ResponseFormat{Type: "json_object"},
		Messages:       msgs,
	}

	content, err := s.provider.Chat(req)
	if err != nil {
		return nil, fmt.Errorf("LLM refine call failed: %w", err)
	}

	// Bersihkan kemungkinan code fence meski sudah dilarang
	cleaned := stripCodeFence(content)

	var result RefineIdeaResult
	if err := json.Unmarshal([]byte(cleaned), &result); err != nil {
		return nil, fmt.Errorf("parse LLM refine response: %w (raw: %s)", err, truncate(cleaned, 200))
	}
	return &result, nil
}

func ifEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

func stripCodeFence(s string) string {
	out := s
	for _, fence := range []string{"```json", "```"} {
		for len(out) >= len(fence) && out[:len(fence)] == fence {
			out = out[len(fence):]
		}
	}
	return out
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// ProviderName returns the name of the active provider
func (s *LLMService) ProviderName() string {
	return s.provider.Name()
}

// ReviewEvidence sends evidence to LLM and returns verdict, rationale, next action
func (s *LLMService) ReviewEvidence(content, evidenceType string) (verdict, rationale, nextAction string, err error) {
	systemPrompt := `Kamu adalah reviewer bukti validasi bisnis F&B. Nilai apakah bukti yang diberikan pengguna cukup valid.

Kriteria:
- valid: bukti jelas, relevan, dan menunjukkan aksi nyata
- weak: bukti ada tapi kurang jelas atau kurang meyakinkan
- invalid: bukti tidak relevan atau tidak bermakna

Output JSON:
{"verdict": "valid|weak|invalid", "rationale": "alasan singkat dalam Bahasa Indonesia", "next_action": "continue|repeat"}`

	userPrompt := fmt.Sprintf(`Tipe bukti: %s\n\nIsi bukti:\n"%s"\n\nBeri penilaian.`, evidenceType, content)

	req := ChatRequest{
		Temperature:    0.2,
		ResponseFormat: &ResponseFormat{Type: "json_object"},
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
	}

	resp, callErr := s.provider.Chat(req)
	if callErr != nil {
		return "", "", "", callErr
	}

	var result struct {
		Verdict    string `json:"verdict"`
		Rationale  string `json:"rationale"`
		NextAction string `json:"next_action"`
	}
	if callErr := json.Unmarshal([]byte(resp), &result); callErr != nil {
		return "", "", "", callErr
	}

	switch result.Verdict {
	case "valid", "weak", "invalid":
	default:
		result.Verdict = "valid"
	}
	if result.NextAction != "repeat" {
		result.NextAction = "continue"
	}
	if result.Rationale == "" {
		result.Rationale = "Bukti diterima dan diverifikasi."
	}

	return result.Verdict, result.Rationale, result.NextAction, nil
}

// ProviderFactory creates the appropriate LLM provider based on config
func ProviderFactory(providerName, baseURL, apiKey, model string) LLMProvider {
	switch providerName {
	case "ollama":
		return NewOllamaProvider("ollama", baseURL, model)
	case "mock":
		return NewMockProvider()
	default:
		// openai, anthropic, opencode, openrouter, groq, together, etc.
		// All use OpenAI-compatible API format with different base URLs
		return NewOpenAIProvider(providerName, baseURL, apiKey, model)
	}
}
