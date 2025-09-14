package prompt_builder

func (pb *PromptBuilder) generateNewPrompt(prompt Prompt) string {
	pb.sanitizer.SetPrompt(prompt.value)
	var sanitizedPrompt string = pb.sanitizer.sanitizePrompt()
	pb.requestStack.Push(sanitizedPrompt)

	return sanitizedPrompt
}
