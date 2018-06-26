package main

import (
	"fmt"

	"github.com/flw-cn/go-fortune"
)

func main() {
	o, err := fortune.Draw(
		fortune.Category("tang300", 50),
		fortune.Category("song100", 50),
	)

	fmt.Printf("fortune: %v\n%v\n", err, o)
}
