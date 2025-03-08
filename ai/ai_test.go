package ai

import "testing"

func TestOllama(t *testing.T) {
	msg := Message{
		Role:    "user",
		Content: "你是什么颜色的",
	}
	req := ChatRequest{
		Model:    "deepseek-r1:1.5b",
		Stream:   false,
		Messages: []Message{msg},
	}

	response, err := sendToOllama("http://localhost:11434/api/chat", req)
	if err != nil {
		panic(err)
	}
	println(response.Message.Content)
}
