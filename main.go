package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func downloadAndExtractTextFiles(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return "", err
	}

	var content strings.Builder

	for _, f := range zipReader.File {
		if (strings.HasSuffix(f.Name, ".md")) && strings.Contains(f.Name, "/docs/") {
			rc, err := f.Open()
			if err != nil {
				return "", err
			}
			defer rc.Close()

			fileContent, err := io.ReadAll(rc)
			if err != nil {
				return "", err
			}

			content.Write(fileContent)
			content.WriteString("\n\n")
		}
	}

	return content.String(), nil
}

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-2.5-flash")

	promptBytes, err := os.ReadFile("prompt.md")
	if err != nil {
		log.Fatal(err)
	}
	command := string(promptBytes)

	documentationContent, err := downloadAndExtractTextFiles("https://github.com/google/adk-docs/archive/refs/heads/main.zip")
	if err != nil {
		log.Fatal(err)
	}

	prompt := command + "\n\n" + documentationContent

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					fmt.Println(txt)
				}
			}
		}
	}
}
