package mapkha

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	m "github.com/veer66/mapkha"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func wordcount_nospace(words []string) int{
	wordcount := 0
	for _, word := range words {
		if word != " " {
			wordcount++
		}
	}
	return wordcount
}

func WordCountNoSpace(words string) int{
	dict, e := m.LoadDefaultDict()
	check(e)
	wordcut := m.NewWordcut(dict)
	cut_words := wordcut.Segment(words)
	fmt.Println(strings.Join(cut_words,"|"))
	return wordcount_nospace(cut_words)
}

func Test() {
    dict, e := m.LoadDefaultDict()
    check(e)
    wordcut := m.NewWordcut(dict)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(strings.Join(wordcut.Segment(scanner.Text()), "|"))
    }
}
