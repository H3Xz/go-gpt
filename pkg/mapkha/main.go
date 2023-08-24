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

func test() {
    dict, e := m.LoadDefaultDict()
    check(e)
    wordcut := m.NewWordcut(dict)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(strings.Join(wordcut.Segment(scanner.Text()), "|"))
    }
}
