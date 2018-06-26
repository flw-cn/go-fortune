package fortune

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

// Draw return a fortune words for caller
//
// The current implemention of fortune just called fortune(6) directly.
// For more friendly, it trim all leading-space and trailing-space,
// and any SGR codes. More about SGR:
// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
func Draw() (string, error) {
	re := regexp.MustCompile("\x1b[^m]*m")

	cmd := exec.Command("fortune")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("os/exec: %s", err)
	}

	output = bytes.Trim(output, " \t\n")
	output = re.ReplaceAll(output, nil)

	return string(output), nil
}

// DrawN call Draw N times and collects result
func DrawN(count int) (result []string, err error) {
	if count <= 0 {
		count = 1
	}

	for i := 0; i < count; i++ {
		output, err := Draw()
		if err != nil {
			return nil, err
		}
		result = append(result, output)
	}

	return
}
