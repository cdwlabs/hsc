package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/mitchellh/cli"
	"github.com/visionmedia/go-spin"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	log.SetOutput(ioutil.Discard)
	//	s := spin.New()
	//	show(s, "Default", spin.Default)

	//	show(s, "Box1", spin.Box1)
	//	show(s, "Box2", spin.Box2)
	//	show(s, "Box3", spin.Box3)
	//	show(s, "Box4", spin.Box4)
	//	show(s, "Box5", spin.Box5)
	//	show(s, "Box6", spin.Box6)
	//	show(s, "Box7", spin.Box7)

	//	show(s, "Spin1", spin.Spin1)
	//	show(s, "Spin2", spin.Spin2)
	//	show(s, "Spin3", spin.Spin3)
	//	show(s, "Spin4", spin.Spin4)
	//	show(s, "Spin5", spin.Spin5)
	//	show(s, "Spin6", spin.Spin6)
	//	show(s, "Spin7", spin.Spin7)
	//	show(s, "Spin8", spin.Spin8)
	//	show(s, "Spin9", spin.Spin9)

	// If command line args include "--version" or "-v"; then just show the version.
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("hsc"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}

func show(s *spin.Spinner, name, frames string) {
	s.Set(frames)
	fmt.Printf("\n\n  %s: %s\n\n", name, frames)
	for i := 0; i < 30; i++ {
		fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}
}
