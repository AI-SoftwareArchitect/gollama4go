package entities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (g *G4GClient) SendPromptRequest(request PromptRequest) (*PromptResponse, error) {
	// 1. JSON'a çevir
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("json marshal failed: %w", err)
	}

	// 2. Endpoint belirle
	endpoint := fmt.Sprintf("%s/api/prompt", g.BaseURL)

	// 3. HTTP POST
	resp, err := g.Http.Post(endpoint, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http post failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	// 4. Status check
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// 5. Decode JSON response
	var pr PromptResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return nil, fmt.Errorf("json decode failed: %w", err)
	}

	return &pr, nil
}
