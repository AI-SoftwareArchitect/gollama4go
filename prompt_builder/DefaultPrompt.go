package prompt_builder

type DefaultPrompt struct {
	Prompt
}

func (d *DefaultPrompt) Build() string {
	return d.value
}
