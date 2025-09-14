package entities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type modelTag struct {
	Name string `json:"name"`
}

type listModelsResponse struct {
	Models []modelTag `json:"models"`
}

func (g *G4GClient) ListOllamaModels() (string, error) {
	resp, err := g.Http.Get(g.BaseURL + "/api/tags")
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama returned %d: %s", resp.StatusCode, string(body))
	}

	var data listModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	names := ""
	for i, m := range data.Models {
		if i > 0 {
			names += ", "
		}
		names += m.Name
	}

	return names, nil
}
