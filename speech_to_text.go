package lana

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os/exec"

	"github.com/mattetti/audio/riff"
)

type response struct {
	Result []Result `json:"result"`
}

type Result struct {
	Alternative []alternative `json:"alternative"`
}

type alternative struct {
	Transcript string  `json:"transcript"`
	Confidence float64 `json:"confidence"`
}

func Listen(lang string) (result *Result, err error) {
	cmd := exec.Command("arecord", "-f", "S16_LE", "-r", "16000")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	buf := new(bytes.Buffer)
	wait(io.TeeReader(stdout, buf))

	// send speech
	resp, err := http.Post(
		fmt.Sprintf("https://www.google.com/speech-api/v2/recognize?output=json&lang=%s&key=AIzaSyBOti4mM-6x9WDnZIjIeyEU21OpBXqWBgw", lang),
		"audio/l16; rate=16000",
		buf,
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	Say("ok.", "en")

	// parse response
	s := bufio.NewScanner(resp.Body)
	s.Scan() // the first line is useless
	s.Scan()

	var r response
	err = json.Unmarshal(s.Bytes(), &r)
	if err != nil {
		return
	}
	if len(r.Result) > 0 {
		// TODO: more experiments
		result = &r.Result[0]
	}
	return
}

// subSampleWindow * sampleWindow is about sample rate
const (
	subSampleWindow = 200
	sampleWindow    = 100
)

func wait(r io.Reader) {
	c := riff.New(r)
	// parse wav header
	err := c.ParseHeaders()
	if err != nil {
		return
	}

	soundData, err := c.NextChunk()
	if err != nil {
		return
	}

	var (
		prev   int16
		sum    int
		count  int
		window []int
	)
	for {
		var next int16
		soundData.ReadLE(&next)

		// measure loudness
		diff := int(next) - int(prev)
		if diff < 0 {
			diff = -diff
		}
		if sum < math.MaxInt64-diff {
			sum += diff
		}
		prev = next
		count++

		// evaluate avg loudness
		if count == subSampleWindow {
			avg := sum / count
			if len(window) < sampleWindow {
				window = append(window, avg)
			} else {
				window = append(window[1:], avg)
				if noSound(window) {
					break
				}
			}

			sum = 0
			count = 0
		}
	}
}

func noSound(window []int) bool {
	var sum int
	for _, n := range window {
		if sum < math.MaxInt64-n {
			sum += n
		}
	}
	avg := sum / sampleWindow
	for _, n := range window {
		if n > 2*avg {
			return false
		}
	}
	return true
}
