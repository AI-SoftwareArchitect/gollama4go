package entities

type PromptResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}
