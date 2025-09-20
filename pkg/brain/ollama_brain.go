package brain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mouuff/GoSubAI/pkg/types"
	"github.com/ollama/ollama/api"
)

type Schema struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}

type Property struct {
	Type string `json:"type"`
}

type OllamaBrain struct {
	Client *api.Client
}

func getClient(hosturl string) (*api.Client, error) {
	if hosturl == "" || hosturl == "default" {
		return api.ClientFromEnvironment()
	} else {
		u, err := url.Parse(hosturl)

		if err != nil {
			return nil, err
		}

		return api.NewClient(u, http.DefaultClient), nil
	}
}

func NewOllamaBrain(hosturl string) (*OllamaBrain, error) {
	client, err := getClient(hosturl)

	if err != nil {
		return nil, err
	}

	return &OllamaBrain{
		Client: client,
	}, nil
}

func (c *OllamaBrain) GenerateString(ctx context.Context, r *types.PromptRequest) (string, error) {
	formatSchema := Schema{
		Type: "object",
		Properties: map[string]Property{
			r.PropertyName: {
				Type: "string",
			},
		},
		Required: []string{r.PropertyName},
	}

	var result map[string]string

	respFunc := func(resp api.GenerateResponse) error {
		err := json.Unmarshal([]byte(resp.Response), &result)
		if err != nil {
			return fmt.Errorf("failed to parse response: %v", err)
		}
		return nil
	}

	err := c.generate(ctx, r, formatSchema, respFunc)

	if err != nil {
		return "", fmt.Errorf("failed to generate: %v", err)
	}

	return result[r.PropertyName], nil
}

func (c *OllamaBrain) generate(ctx context.Context, r *types.PromptRequest, formatSchema Schema, fn api.GenerateResponseFunc) error {
	format, err := json.Marshal(formatSchema)
	if err != nil {
		return fmt.Errorf("failed to marshal the format schema: %v", err)
	}

	req := &api.GenerateRequest{
		Model:  r.Model,
		Prompt: r.Prompt,
		System: r.SystemPrompt,
		Format: format,

		// set streaming to false
		Stream: new(bool),
	}

	err = c.Client.Generate(ctx, req, fn)
	if err != nil {
		return fmt.Errorf("failed to generate response: %v", err)
	}

	return nil
}
