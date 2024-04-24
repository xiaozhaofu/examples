// @Author xiaozhaofu 2023/4/28 16:33:00
package repository

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/E"
	"github.com/gtkit/encry/md5"
	"github.com/gtkit/goerr"
	"github.com/sashabaranov/go-openai"

	"ydsd_chat/config"
)

var ChatMsg = make(chan string)

func (r *reposit) sendStremGpt(ccp *gin.Context, chatargs *chatUserArgs, key string) (openai.ChatCompletionStreamResponse, error) {
	var baseurl string

	msg := r.chatmessage(chatargs)

	// 根据key来使用不同的 baseurl
	baseurl = "http://chat1.officeaid02.com/v1"
	if E.InSlice(config.PaidChatKey(), key) {
		baseurl = "https://chat.officeaid02.com/v1"
	}

	ctx, cancel := context.WithTimeout(ccp.Request.Context(), 120*time.Second)
	defer cancel()

	openConfig := openai.DefaultConfig(key)
	openConfig.BaseURL = baseurl
	openConfig.HTTPClient = &http.Client{
		Timeout: 120 * time.Second,
	}

	u := md5.New(strconv.Itoa(int(chatargs.uid)) + "ydsdassistant")

	// 开始计算请求时间
	// sTime := time.Now()
	c := openai.NewClientWithConfig(openConfig)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		// MaxTokens: 20,
		Messages: msg,
		User:     u,
		Stream:   true,
	}

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		// return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if goerr.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return openai.ChatCompletionStreamResponse{}, nil
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			// return
		}

		// 计算请求使用时间
		// subTime := time.Since(sTime)
		// fmt.Printf("---- response.Choices Content: %s---chatGpt响应时间:%+v\n", response.Choices[0].Delta.Content, subTime)
		ChatMsg <- response.Choices[0].Delta.Content

	}
	// return openai.ChatCompletionResponse{}, nil

}
