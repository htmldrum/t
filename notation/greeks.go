package notation

import (
	"fmt"
	"os"
)

type Greeks []string

func Not(){
	fmt.Fprintf(os.Stdout, "Wut\n", )
	fmt.Println("Where?")
}

func PUsage() string {
	return "not - Prints common mathematical operators and their description\n"
}

func PCommands() []string {
	var com []string
	com = append(com, "not")
	return com
}
