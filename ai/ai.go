package ai

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   Message   `json:"message"`
	Done      bool      `json:"done"`
}

func sendToOllama(url string, req ChatRequest) (*ChatResponse, error) {
	payload, _ := json.Marshal(req)
	client := &http.Client{}
	httpReq, _ := http.NewRequest("POST", url, bytes.NewReader(payload))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ollamaResp ChatResponse
	json.NewDecoder(resp.Body).Decode(&ollamaResp)
	return &ollamaResp, nil
}
