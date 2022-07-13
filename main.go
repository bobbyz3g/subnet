package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: ./subnet <subnet string>")
		os.Exit(0)
	}
	sn := args[1]

	// TODO(Kaiser925): add handle logic
	fmt.Println(sn)
}
