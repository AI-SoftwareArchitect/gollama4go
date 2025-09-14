package prompt_builder

type PromptType int

const (
	Default = iota
	RoleBased
	NegativePromptType
	TreeOfThought
	Chain
	ReAct
	ChainOfThought
)
