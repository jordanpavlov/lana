package lana

import "testing"

func TestListen(t *testing.T) {
	t.Log(Listen("zh_TW"))
}

func TestSay(t *testing.T) {
	Say("I am happy", "en")
	Say("我很快樂", "zh_TW")
}
