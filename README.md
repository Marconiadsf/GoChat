# GoChat

A provider-agnostic CLI chatbot written in Go.

Built as a Go learning project, exploring idiomatic patterns and standard library usage. Supports Claude, Gemini, and Ollama out of the box — swap providers by editing a single config file.

## What this covers

- Implicit interfaces (`Provider` satisfied by `Claude`, `Gemini`, `Ollama`)
- Dependency inversion — `llm` defines the contract; `main` wires concrete adapters (composition root / Ports & Adapters)
- Structs and methods
- Idiomatic error handling
- Native HTTP client (`net/http`)
- JSON encoding/decoding (`encoding/json`)
- YAML config parsing (`gopkg.in/yaml.v3`)
- Package organization

## Setup

**1. Clone and enter the project:**

```bash
git clone https://github.com/Marconiadsf/GoChat.git
cd GoChat
```

**2. Create your config file from the example:**

```bash
cp config.example.yaml config.yaml
```

**3. Edit `config.yaml` with your provider and API key:**

```yaml
provider: gemini          # claude | gemini | ollama
api_key: YOUR_KEY_HERE
model: gemini-2.5-flash
```

Where to get API keys:
- **Claude** → [console.anthropic.com](https://console.anthropic.com/)
- **Gemini** → [aistudio.google.com](https://aistudio.google.com/app/apikey)
- **Ollama** → no key needed, just leave `api_key` blank and [install Ollama](https://ollama.com)

**4. Run:**

```bash
go run main.go
```

## Usage

```
────────────────────────────────────────────────────
  GoChat  •  gemini  •  gemini-2.5-flash
────────────────────────────────────────────────────

Você › hello!

────────────────────────────────────────────────────
gemini
────────────────────────────────────────────────────
Hello! How can I help you today?
────────────────────────────────────────────────────

Você › 
```

The chat keeps conversation history in memory for the duration of the session — the full message history is sent to the provider on each turn, enabling multi-turn context.

## Switching providers

Edit `config.yaml` and restart. No code changes needed — that's the point of the `Provider` interface.

```yaml
# Switch to local Ollama
provider: ollama
api_key: ""
model: llama3.2
```

```yaml
# Switch to Claude
provider: claude
api_key: sk-ant-...
model: claude-haiku-4-5-20251001
```
