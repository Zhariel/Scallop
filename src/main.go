package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var history []string
	env := loadEnv()
	reader := bufio.NewReader(os.Stdin)
	// w := createWindow(900, 500, "Scallop")
	// w.ShowAndRun()

	for {
		fmt.Println(">")
		handleInput(env, history, reader)

		break
	}
}
