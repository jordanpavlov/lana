package lana

import "testing"

func TestListen(t *testing.T) {
	r, err := Listen("en")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(r)
	}
}

func TestSay(t *testing.T) {
	Say("I am happy", "en")
	Say("我很快樂", "zh_TW")
}
