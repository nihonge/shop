package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// 测试函数
func TestAuthEndpoint(t *testing.T) {
	// 1. 定义请求体
	requestBody := map[string]any{"user_id": 2}
	jsonBody, _ := json.Marshal(requestBody)

	// 2. 创建 HTTP 请求对象
	req, err := http.NewRequest("POST", "http://localhost:8888/auth", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("创建请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 3. 发送请求（使用默认 HTTP 客户端）
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("请求发送失败: %v", err)
	}
	defer resp.Body.Close()

	// 4. 验证状态码
	if resp.StatusCode != http.StatusOK {
		t.Errorf("期望状态码 200，实际状态码: %d", resp.StatusCode)
	}

	// 5. 解析响应体
	var responseBody map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Errorf("解析响应体失败: %v", err)
	}
	// 6. 打印响应体，以JSON格式
	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		t.Errorf("格式化json失败: %v", err)
	}
	fmt.Println("返回的JSON数据:")
	fmt.Println(string(formattedJSON))
}
