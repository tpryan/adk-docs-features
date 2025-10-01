# ADK Docs Feature Generator

This project contains a Go application and a GitHub Action workflow designed to automatically generate documentation for the `google/adk-docs` repository. It uses the Google Gemini API to create new content based on the existing documentation and a predefined prompt.

## How it Works

The core of this project is a Go application that:
1.  Fetches a zip archive of the `main` branch of the `google/adk-docs` repository.
2.  Extracts the content of all markdown files found in the `/docs/` directory.
3.  Combines the extracted documentation with a base prompt (`prompt.md`) and an example of the desired output (`example.md`).
4.  Sends the complete prompt to the Google Gemini API (`gemini-1.5-pro` model).
5.  Saves the generated markdown content to a file named `output_1.md`.

## GitHub Action Workflow

The primary intended use of this project is via the included GitHub Action. This workflow is designed to be placed in the `adk-docs` repository. When triggered (e.g., by a push to `main`), it automates the process of generating and proposing new documentation.

The workflow performs the following steps:
1.  **Checkout `adk-docs`:** Checks out the code of the repository where the action is running.
2.  **Checkout `adk-docs-features`:** Checks out the source code from this repository (`tpryan/adk-docs-features`) into a subdirectory.
3.  **Setup Go:** Initializes the Go environment.
4.  **Run Program:** Executes the Go application, which generates the `output_1.md` file.
5.  **Prepare PR Content:** Moves the generated file to `docs/features/index.md` within the `adk-docs` repository structure.
6.  **Create Pull Request:** Automatically creates a new branch and opens a pull request with the newly generated `docs/features/index.md` file.

### Usage

1.  **Add the workflow:** Copy the `.github/workflows/run.yml` file from this repository to the target repository (e.g., `adk-docs/.github/workflows/run.yml`).
2.  **Set secrets:** In the GitHub settings for the `adk-docs` repository, add a new "Repository secret" named `GEMINI_API_KEY`. The value should be your API key for the Gemini API.
3.  The action will run automatically on the next push to the `main` branch.

## Local Development

You can run the Go application locally to test changes or generate content manually.

### Prerequisites
*   Go (version 1.22 or newer)
*   A Gemini API Key

### Steps
1.  **Install dependencies:**
    ```sh
    go mod tidy
    ```
2.  **Set the environment variable:**
    ```sh
    export GEMINI_API_KEY="YOUR_API_KEY"
    ```
3.  **Run the application:**
    ```sh
    go run main.go
    ```
4.  The output will be written to `output_1.md` in the project root.

## Configuration

The generated content is primarily controlled by the prompt files:
*   `prompt.md`: Contains the main instructions and queries sent to the Gemini model.
*   `example.md`: Provides a detailed, high-quality example of the desired output format and content, which helps the model produce a consistent and accurate result.

To change the generated output, modify these files.
