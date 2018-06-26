package fortune

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

type options struct {
	category string
	precent  int
}

func Category(name string, precent int) options {
	return options{
		category: name,
		precent:  precent,
	}
}

// Draw return a fortune words for caller
//
// The current implemention of fortune just called fortune(6) directly.
// For more friendly, it trim all leading-space and trailing-space,
// and any SGR codes. More about SGR:
// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
func Draw(opts ...options) (string, error) {
	re := regexp.MustCompile("\x1b[^m]*m")

	sum := 0
	for _, o := range opts {
		sum += o.precent
	}

	args := []string{}
	for _, o := range opts {
		p := 0
		if sum > 0 {
			p = int(float64(o.precent) / float64(sum) * 100)
		} else {
			p = int(1.0 / float64(len(opts)) * 100)
		}
		args = append(args, fmt.Sprintf("%d%%", p), o.category)
	}

	fmt.Printf("args: %#v\n", args)
	cmd := exec.Command("fortune", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("os/exec: %s", err)
	}

	output = bytes.Trim(output, " \t\n")
	output = re.ReplaceAll(output, nil)

	return string(output), nil
}

// DrawN call Draw N times and collects result
func DrawN(count int, opts ...options) (result []string, err error) {
	if count <= 0 {
		count = 1
	}

	for i := 0; i < count; i++ {
		output, err := Draw(opts...)
		if err != nil {
			return nil, err
		}
		result = append(result, output)
	}

	return
}
