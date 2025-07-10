package main

import (
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type ToolCallBuilder struct {
	ID        string
	Name      string
	Arguments strings.Builder
}

func main() {
	stream := fakeOpenAIStream()
	toolCalls := make(map[int]*ToolCallBuilder)

	for _, resp := range stream { // 多帧

		for _, choice := range resp.Choices {

			for _, tool := range choice.Delta.ToolCalls {
				builder, exists := toolCalls[*tool.Index]
				if !exists {
					builder = &ToolCallBuilder{
						ID:   tool.ID,
						Name: tool.Function.Name,
					}
					toolCalls[*tool.Index] = builder
				}
				builder.Arguments.WriteString(tool.Function.Arguments)
			}
		}
	}

	// 打印组合完成的 ToolCall
	for _, builder := range toolCalls {
		fmt.Println("== Tool Call ==")
		fmt.Println("ID:", builder.ID)
		fmt.Println("Name:", builder.Name)
		fmt.Println("Arguments:", builder.Arguments.String())
	}
}

func fakeOpenAIStream() []openai.ChatCompletionStreamResponse {
	index := 0
	return []openai.ChatCompletionStreamResponse{
		{
			Choices: []openai.ChatCompletionStreamChoice{
				{
					Delta: openai.ChatCompletionStreamChoiceDelta{
						ToolCalls: []openai.ToolCall{
							{
								Index: &index,
								ID:    "toolcall-123",
								Function: openai.FunctionCall{
									Name:      "get_weather",
									Arguments: "{",
								},
							},
						},
					},
				},
			},
		},
		{
			Choices: []openai.ChatCompletionStreamChoice{
				{
					Delta: openai.ChatCompletionStreamChoiceDelta{
						ToolCalls: []openai.ToolCall{
							{
								Index: &index,
								Function: openai.FunctionCall{
									Arguments: "\"city\":\"Beijing\"}",
								},
							},
						},
					},
				},
			},
		},
	}
}
