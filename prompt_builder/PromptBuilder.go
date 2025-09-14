package prompt_builder

type PromptBuilder struct {
	promptType   PromptType
	requestStack *PromptStack
	sanitizer    *Sanitizer
}
