package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const PROMPT = "$ "

type Shell struct {
	builtins map[string]func(args []string) error
	wd       string
}

func main() {
	sh := &Shell{
		builtins: make(map[string]func(args []string) error),
		wd:       "/app",
	}

	sh.builtins["exit"] = sh.exit

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(PROMPT)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		if cmd, ok := sh.builtins[args[0]]; ok {
			if err := cmd(args); err != nil {
				fmt.Println(err)
			}
			continue
		}

		fmt.Printf("%s: command not found\n", args[0])
	}
}

func (sh *Shell) exit(args []string) error {
	os.Exit(0)
	return nil
}
