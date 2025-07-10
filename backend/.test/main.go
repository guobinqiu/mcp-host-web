package main

import (
	"fmt"
	"strings"
)

func main() {

	ch := make(chan string)

	s := `function add(a,b) {return a+b;} function sub(a,b) {return a-b;}`

	tokens := strings.Split(s, " ")

	ss := []string{}

	var sb strings.Builder
	go func(ch <-chan string) {
		for token := range ch {
			fmt.Println(token)

			if token == "function" {
				ss = append(ss, sb.String())
				sb.Reset()
			}
			sb.WriteString(token)
			sb.WriteByte(' ')
			//time.Sleep(1 * time.Second)
		}
	}(ch)

	for _, token := range tokens {
		ch <- token
	}

	close(ch)

	// fmt.Println("===", sb.String())
	ss = append(ss, sb.String())

	for _, s := range ss {
		fmt.Println(s)
	}
}
