package lana

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

const (
	limit = 100
)

func Say(text, lang string) {
	for _, s := range prepare(text) {
		exec.Command("mpg123", fmt.Sprintf(
			"http://translate.google.com/translate_tts?ie=UTF-8&tl=%s&q=%s",
			lang,
			url.QueryEscape(s),
		)).Run()
	}
}

func prepare(s string) (sentences []string) {
	// token must be under limit
	var tokens []string
	for _, chunk := range strings.Fields(s) {
		if chunk == "" {
			continue
		}
		runes := []rune(chunk)
		if len(runes) > limit {
			for i := 0; i < len(runes); i += limit {
				end := i + limit
				if end > len(runes) {
					end = len(runes)
				}
				tokens = append(tokens, string(runes[i:end]))
			}
		} else {
			tokens = append(tokens, chunk)
		}
	}
	for _, token := range tokens {
		if len(sentences) == 0 ||
			len([]rune(sentences[len(sentences)-1]))+len([]rune(token))+1 > limit {
			sentences = append(sentences, token)
		} else {
			sentences[len(sentences)-1] += " " + token
		}
	}
	return
}
