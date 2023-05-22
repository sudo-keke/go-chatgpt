package common

import (
	"github.com/sashabaranov/go-openai"
)

func connect() *openai.Client {
	return openai.NewClient("")
}
