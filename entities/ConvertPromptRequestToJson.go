package entities

import (
	"encoding/json"
)

func convertPromptRequestToJson(request PromptRequest) (string, string) {
	data, err := json.Marshal(request)
	if err != nil {
		return "", "dequeue"
	}
	return string(data), "success"
}
