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

func buildProvider(cfg config.Config) (llm.Provider, error) {
	switch cfg.Provider {
	case "claude":
		return llm.NewClaude(cfg.APIKey, cfg.Model), nil
	case "gemini":
		return llm.NewGemini(cfg.APIKey, cfg.Model), nil
	case "ollama":
		return llm.NewOllama(cfg.Model, "http://localhost:11434"), nil
	default:
		return nil, fmt.Errorf("provider desconhecido: %s", cfg.Provider)
	}
}

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	provider, err := buildProvider(cfg)
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
