package content_generate

import chatgpt "command-line-arguments/Users/s4intz/Desktop/Project/Golang/go-gpt/pkg/chatgpt/main.go"

func GetArticle(keyword string) {
	
	client := chatgpt.GetClient()
	chatgpt.PromptContentGenerator(client,keyword)
	
}