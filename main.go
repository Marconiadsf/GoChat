package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Marconiadsf/GoChat/chat"
	"github.com/Marconiadsf/GoChat/config"
	"github.com/Marconiadsf/GoChat/llm"
)

const (
	reset = "\033[0m"
	bold  = "\033[1m"
	cyan  = "\033[96m"
	green = "\033[92m"
	gray  = "\033[90m"
)

func line() {
	fmt.Println(gray + strings.Repeat("─", 52) + reset)
}

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	provider, err := llm.NewProvider(cfg.Provider, cfg.APIKey, cfg.Model)
	if err != nil {
		log.Fatal(err)
	}

	c := chat.NewChat(provider)

	fmt.Println()
	line()
	fmt.Printf("  GoChat  •  %s  •  %s\n", cfg.Provider, cfg.Model)
	line()
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(cyan + bold + "Você" + reset + gray + " › " + reset)
		if !scanner.Scan() {
			break
		}
		message := strings.TrimSpace(scanner.Text())
		if message == "" {
			continue
		}

		fmt.Println()
		response, err := c.Send(message)
		if err != nil {
			log.Fatal(err)
		}

		line()
		fmt.Println(green + bold + cfg.Provider + reset)
		line()
		fmt.Println(response)
		line()
		fmt.Println()
	}
}
