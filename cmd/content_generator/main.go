package content_generator

import "githun.com/H3Xz/go-gpt/pkg/chatgpt"

func GetArticle(keyword string) {
	
	client := chatgpt.GetClient()
	chatgpt.PromptContentGenerator(client,keyword)
	
}