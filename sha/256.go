package sha

import (
	"fmt"
	"bytes"
)

func Sha() {
	var buf bytes.Buffer
	for i := 0; i < 5; i++ {
		buf.WriteString("Computing input")
		for j := 0; j < 3 % i; j++ {
			buf.WriteString(".")
		}
		fmt.Printf(buf.String())
	}
}

func PUsage() string {
	return "sha - computes the sha256 hash of input \n"
}

func PCommands() []string {
	var com []string
	com = append(com, "sha")
	return com
}
