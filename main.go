package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-1.5-flash")

	command := `
Analyze the documentation for the adk project and please come up with a list of features that adk delivers. Focus on page titles and section headers to come up with that list.

Once that is done, I'd like to generate a matrix of features enabled by runtime.

Currently there are two runtimes python which is the primary, and java. But soon go and typescript will be added, but typescript will be known as adk-js.

A good metric for figuring out if a feature is supported per runtime is if the sample code associated with a given feature in adk-docs includes samples in a given language.

Ideally the output format would be markdown
`

	// In a real script, you would load the content of your documentation files here.
	// For this example, we'll use a placeholder.
	documentationContent := "PASTE_ALL_YOUR_DOCUMENTATION_CONTENT_HERE"

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
