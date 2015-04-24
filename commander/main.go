package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/SlyMarbo/rss"
	"github.com/taylorchu/lana"
)

func say(s string) {
	fmt.Println(s)
	lana.Say(s, "en")
}

func listen() string {
	r, err := lana.Listen("en")
	if err != nil || len(r.Alternative) == 0 {
		return ""
	}
	s := strings.ToLower(r.Alternative[0].Transcript)
	if s != "" {
		say("ok")
		fmt.Println(s)
	}
	return s
}

func main() {
	say("What can I do for you?")
	cmd := listen()
	switch {
	case regexp.MustCompile("^google ").MatchString(cmd):
		say("Searching the result.")
		exec.Command("xdg-open",
			fmt.Sprintf("https://www.google.com/#q=%s", url.QueryEscape(strings.TrimPrefix(cmd, "google "))),
		).Run()
	case regexp.MustCompile("^say ").MatchString(cmd):
		say(strings.TrimPrefix(cmd, "say "))
	case regexp.MustCompile("gmail").MatchString(cmd):
		say("Checking your email.")
		exec.Command("xdg-open", "https://gmail.com").Run()
	case regexp.MustCompile("time").MatchString(cmd):
		say(fmt.Sprintf("The time is %s", time.Now().Format("Jan 2, 2006 at 3:04pm")))
	case regexp.MustCompile("what.*your name").MatchString(cmd):
		say("My name is lana.")
	case regexp.MustCompile("news").MatchString(cmd):
		say("Checking bbc news.")
		feed, err := rss.Fetch("http://feeds.bbci.co.uk/news/rss.xml")
		if err != nil {
			return
		}
		for _, item := range feed.Items {
			say(item.Content)
			if "stop" == listen() {
				break
			}
		}
	default:
		say("Please try again.")
	}
}
