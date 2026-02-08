// Package main is the main entry point of the application.
// It handles command-line arguments and routes execution to specific modules.
//
// The application supports modular execution where each module performs
// distinct operations. Command-line arguments determine which module to run.
//
// Usage:
//
//	go run main.go [module] [options]
//
// Available modules:
//   - test    : Run test
//
// Examples:
//
//	go run main.go test
//	make run arg=test
//
// For detailed help on a specific module:
//
//	go run main.go [module] --help
package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hello, World!", os.Args[1])
}
