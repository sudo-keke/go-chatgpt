package common

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

const (
	// MODEL 启用模型，默认3.5，目前不支持GPT4
	MODEL = openai.GPT3Dot5Turbo
)

// ChatText 文字类型信息,只能单句话聊天，没有上下文语境
func ChatText(content string) (string, error) {
	client := connect()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("chatGPT 响应错误: %v\n", err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

// AudioToText 		音频语音转文本,默认自动识别语言，如需中文翻译请指定Language
// FilePath		:	音频文件地址
func AudioToText(FilePath, Language string) (string, error) {
	client := connect()
	ctx := context.Background()

	/*
		req := AudioRequest{
					FilePath:    path,
					Model:       "whisper-3",
					Prompt:      "用简体中文",
					Temperature: 0.5,
					Language:    "zh",
				}
	*/
	req := openai.AudioRequest{
		// 音频模型 whisper-1
		Model:    openai.Whisper1,
		FilePath: FilePath,
		Language: Language,
	}
	resp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return "", err
	}
	return resp.Text, nil
}

// DallE2Link 生成图像，返回url
// describe 	: 	描述
// size     	: 	图像大小，格式"256x256"，不传默认大小"256x256"
// []byte		：	这里只返回[]byte数组，由用户自定义处理
func DallE2Link(describe, size string) (string, error) {
	client := connect()
	ctx := context.Background()

	if size == "" {
		size = openai.CreateImageSize256x256
	}
	reqUrl := openai.ImageRequest{
		Prompt:         describe,
		Size:           size,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := client.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("图像创建错误: %v\n", err)
		return "", err
	}
	return respUrl.Data[0].URL, nil
}

// DallE2Base64 	生成图像，返回 []byte数组
// describe 	: 	描述
// size     	: 	图像大小，格式"256x256"，不传默认大小"256x256"
// []byte		：	这里只返回[]byte数组，由用户自定义处理
func DallE2Base64(describe, size string) ([]byte, error) {
	client := connect()
	ctx := context.Background()

	if size == "" {
		size = openai.CreateImageSize256x256
	}

	reqBase64 := openai.ImageRequest{
		Prompt:         describe,
		Size:           size,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := client.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("图像创建错误: %v\n", err)
		return nil, err
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 解码错误: %v\n", err)
		return nil, err
	}
	return imgBytes, nil

	/*
		// 后续可直接复制该部分代码使用
		r := bytes.NewReader(imgBytes)
		imgData, err := png.Decode(r)
		if err != nil {
			fmt.Printf("PNG decode error: %v\n", err)
			return
		}

		file, err := os.Create("example.png")
		if err != nil {
			fmt.Printf("File creation error: %v\n", err)
			return
		}
		defer file.Close()

		if err := png.Encode(file, imgData); err != nil {
			fmt.Printf("PNG encode error: %v\n", err)
			return
		}

		fmt.Println("图片保存为example.png")
	*/
}

// ConsoleContextText 控制台根据上下文语境回答，输入 exit 退出
func ConsoleContextText() {
	client := connect()
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)

	// 创建一个新的颜色对象
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	// 打印信息，其中 "对话开始，输入 exit 结束对话" 会被着色为绿色
	//fmt.Println("###############  ", green("对话开始，输入 exit 结束对话"), "  ###############")

	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			fmt.Println(green("\n###############  对话结束，返回上一层  ###############"))
			return
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    MODEL,
				Messages: messages,
			},
		)

		if err != nil {
			fmt.Printf("聊天响应错误: %v\n", err)
			continue
		}

		content := resp.Choices[0].Message.Content
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})

		fmt.Println(cyan(content))
		fmt.Println("---------------------")
	}
}

// ContextText 根据上下文语境回答(同网页版)
func ContextText(text string) (string, error) {
	client := connect()
	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    MODEL,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("聊天响应错误: %v\n", err)
		return "", err
	}

	content := resp.Choices[0].Message.Content
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	return content, nil
}
