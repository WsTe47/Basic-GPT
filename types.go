package main

type CompletionRequest struct {
	Model    Model     `json:"model"`    // Model specifies the GPT model to be used for generating the completion.
	Messages []Message `json:"messages"` // Messages contains the conversation history.
}

type CompletionResponse struct {
	Id      string   `json:"id"`      // Id is a unique identifier for the response.
	Object  string   `json:"object"`  // Object indicates the type of the response.
	Created int      `json:"created"` // Created is a timestamp of when the response was generated.
	Model   Model    `json:"model"`   // Model specifies which GPT model was used.
	Choices []Choice `json:"choices"` // Choices contains the generated outputs.
}
type Choice struct {
	Index        int     `json:"index"`         // Index is the position of the choice in the list.
	Message      Message `json:"message"`       // Message contains the generated output.
	FinishReason string  `json:"finish_reason"` // FinishReason indicates why the generation was stopped.
}

type Message struct {
	Role    Role   `json:"role"`    // Role indicates who sent the message (e.g., user, bot).
	Content string `json:"content"` // Content contains the text of the message.
}
type Model string

type Role string

const (
	RoleUser Role = "user"
)
