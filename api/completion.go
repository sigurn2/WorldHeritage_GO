package api

import (
	"context"
	"fmt"
	"github.com/sigurn2/WorldHeritage_GO/data"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

var client *openai.Client

// init client
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("can't load gpt auth token:%w", err))
	}
	token := viper.Get("token").(string)
	config := openai.DefaultConfig(token)
	config.BaseURL = "https://api.chatanywhere.com.cn/v1"
	client = openai.NewClientWithConfig(config)
}

func getCompletion(q string) string {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: q,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	return resp.Choices[0].Message.Content
}

func GetHeritageOutput(attr data.Attribute, heritage string) string {
	content := fmt.Sprintf("You are a expert on world heritage.Your task is to analyze a world heritage: %v, and choose a option from %v which can describe this heritage best. CONSTRAINT: your answer should only be one of my give options ", heritage, attr.Values)
	return getCompletion(content)
}

// Merge 从 raw data 提取抽象概念
func Merge(attrs []string) {
	content := fmt.Sprintf("I will give you a list, please summarize these values and divide them into some abstract types %v", attrs)
	ans := getCompletion(content)
	fmt.Println(ans)
}

// Select 把 抽象概念 重分类， 得到不相容的属性， 之后会添加到 Attribute 里重新分类
func Select(attrs []string, heritage string) string {
	var content = fmt.Sprintf("Your task is to analyze the attributes of world heritage.Heritage: %vClassify task: Please choose one optionwhich can describe the world cultural heritage best from%v Output limitations:1. Return the most concise answer less than 5 words ", heritage, attrs)
	ans := getCompletion(content)
	return ans
}
