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
    for {
	//say("What can I do for you?")
	fmt.Println("What can I do for you?")    
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
        case regexp.MustCompile("evolution").MatchString(cmd):
		say("Openning evolution")
		exec.Command("evolution").Start()
	case regexp.MustCompile("thunderbird").MatchString(cmd):
		say("Openning thunderbird")
		exec.Command("thunderbird").Start()
	case regexp.MustCompile("g edit").MatchString(cmd): // does not like "G edit"
		say("Openning g edit")
		exec.Command("gedit").Start()
	case regexp.MustCompile("who are you").MatchString(cmd):
		say("I am Lana, nice to meet you!")
	case regexp.MustCompile("chrome").MatchString(cmd):
		say("Openning Chrome")
		exec.Command("google-chrome-stable").Start()
	// Desktop Controll:
	case regexp.MustCompile("quit").MatchString(cmd): // use xdotool, "close" hard to say
		say("Closing")
		exec.Command("xdotool", "key", "Alt_L+F4").Start()
	case regexp.MustCompile("start").MatchString(cmd): // use xdotool
		say("Start")
		exec.Command("xdotool", "key", "Alt_L+F1").Start()
	case regexp.MustCompile("exit").MatchString(cmd):
		say("Leaving")
		return                  // Holy shit GO :)  or break Loop with definend loopname
	default:
		//say("Please try again.")
		fmt.Println("Please try again.")
	}
    }
}
