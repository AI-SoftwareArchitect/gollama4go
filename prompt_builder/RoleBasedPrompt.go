package prompt_builder

type RoleBasedPrompt struct {
	Prompt
	role string
}

func (rb *RoleBasedPrompt) Build() string {
	return "You are " + rb.role + rb.value
}
