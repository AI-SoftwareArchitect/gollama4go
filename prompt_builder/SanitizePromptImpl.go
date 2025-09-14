package prompt_builder

import (
	"regexp"
	"strings"
)

// Sanitizer struct
type Sanitizer struct {
	prompt string
}

// NewSanitizer constructor
func NewSanitizer(p string) *Sanitizer {
	return &Sanitizer{prompt: p}
}

// GetPrompt returns the original prompt
func (s *Sanitizer) GetPrompt() string {
	return s.prompt
}

func (s *Sanitizer) SetPrompt(newPrompt string) {
	s.prompt = newPrompt
}

func (s *Sanitizer) sanitizePrompt() string {
	prompt := s.GetPrompt()
	prompt = strings.TrimSpace(prompt)                               // baş/son boşlukları temizle
	prompt = strings.ReplaceAll(prompt, "\n", " ")                   // newline => space
	prompt = regexp.MustCompile(`\s+`).ReplaceAllString(prompt, " ") // fazla boşlukları sil
	return prompt
}
