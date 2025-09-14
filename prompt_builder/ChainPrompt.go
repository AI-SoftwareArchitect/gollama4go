package prompt_builder

type ChainPrompt struct {
	Prompt
	next *ChainPrompt
}

func (c *ChainPrompt) Build() string {
	var value string = c.Prompt.value
	for c.next != nil {
		value += c.next.value
	}
	return value
}
