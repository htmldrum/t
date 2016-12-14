package main

import (
	"flag"
	"fmt"
	"os"
	"htmldrum.xyz/t/notation"
	"htmldrum.xyz/t/sha"
)

var getArgs = func() []string {
	flag.Parse()
	return flag.Args()
}

func main() {
	args := getArgs()
	if len(args) == 0 || !argInCommands(args[0]){
		printUsageAndExit()
	}
	handleCommand(args[0])
	os.Exit(0)
}

func argInCommands(comm string) bool {
	pCom := gatherCommands()
	for _, c := range *pCom {
		if comm == c {
			return true
		}
	}
	return false
}

func gatherCommands() *[]string{
	var coms []string
	for _, c := range sha.PCommands() {
		coms = append(coms, c)
	}
	for _, c := range notation.PCommands() {
		coms = append(coms, c)
	}
	return &coms
}

func printUsageAndExit() {
	fmt.Printf("t [COMMAND]. Where command is one of the following:\n")
	for _, l := range usage() {
		fmt.Printf(l)
	}
	os.Exit(1)
}

func handleCommand(c string) {
	switch c {
		case "not":
		  notation.Not()
		case "sha":
		  sha.Sha()
	}
}

func usage() (s []string) {
	s = append(s, notation.PUsage())
	s = append(s, sha.PUsage())
	return s
}
