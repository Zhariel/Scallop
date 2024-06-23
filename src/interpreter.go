package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadEnv() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		env[pair[0]] = pair[1]
	}
	return env
}

func parse(history []string, reader *bufio.Reader) []string {
	line, _ := reader.ReadString('\n')
	history = append(history, line)
	return strings.Split(strings.TrimSpace(line), " ")
}

func handleInput(env map[string]string, history []string, reader *bufio.Reader){
	command := parse(history, reader)
	val, ok := env[command[0]]

	if !ok{
		fmt.Print("Command '",command[0],"' recognized.\n")
		return
	}
	fmt.Print(val)
}
