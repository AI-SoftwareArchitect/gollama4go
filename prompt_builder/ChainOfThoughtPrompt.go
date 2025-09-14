package prompt_builder

type ChainOfThoughtPrompt struct {
	Prompt
	headerPrompt string
}

func (c *ChainOfThoughtPrompt) Build() string {
	c.headerPrompt = "First plan and Think step by step:"
	return c.headerPrompt + c.Prompt.value
}
