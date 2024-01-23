package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	apiKey  string
	baseURL string
}

// 这里需要填写的是你的APIkey和url
// url的话就是诸如openai.com。
// 由于众所周知的原因，使用了国内的GPT镜像api
// 可以去此网址自行注册https://referer.shadowai.xyz/r/26448
// 1块钱就可以充值好像，按量计费。
func NewGPTClient() *Client {
	return &Client{
		apiKey:  "APIKey",
		baseURL: "GPTUrl",
	}
}

func (s *Client) SendMsg(question string) (string, error) {
	request := CompletionRequest{
		Messages: []Message{
			{
				Role:    RoleUser,
				Content: question,
			},
		},
		Model: "gpt-3.5-turbo",
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	url := s.baseURL + "/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response CompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
