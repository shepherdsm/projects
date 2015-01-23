package codeabbey

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Bot struct {
	url     string
	body    map[string]string
	content string
}

func NewBot(url, content string) *Bot {
	return &Bot{url, make(map[string]string), content}
}

func (b *Bot) AddToken(token string) {
	b.SetKeyValue("token", token)
}

func (b *Bot) SetKeyValue(k, v string) {
	b.body[k] = v
}

func (b *Bot) RemoveKeyValue(k string) {
	delete(b.body, k)
}

func (b *Bot) MakeRequest() map[string]string {
	reqBody := mapToString(b.body)

	resp, err := http.Post(b.url, b.content, strings.NewReader(reqBody))
	if err != nil {
		log.Fatal("Error processing request: ", err)
	}
	defer resp.Body.Close()

	return respToMap(resp.Body)
}

func respToMap(rd io.Reader) map[string]string {
	keys := make(map[string]string)
	scanner := bufio.NewScanner(rd)

	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ": ")
		keys[items[0]] = items[1]
	}

	return keys
}

func mapToString(m map[string]string) string {
	var out string

	for k, v := range m {
		out += fmt.Sprintf("%s: %s\n", k, v)
	}

	return out
}
