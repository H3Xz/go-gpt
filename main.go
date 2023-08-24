package main

import (
	"log"
	"strconv"

	"github.com/H3Xz/go-gpt/pkg/mapkha"
)

func main(){
	
	log.Println("Running...")

	//content_generator.GetArticle("ขั้นตอนการทำวีซ่าเริ่มต้นยังไง?")

	//mapkha.Test()

	word_count := mapkha.WordCountNoSpace("ขั้นตอนการทำวีซ่าเริ่มต้นยังไง")
	log.Println("จำนวนคำ :"+ strconv.Itoa(word_count))
}