
# GoSubAI
GoSubAI is an AI-powered subtitle translator built in Go.

You can define custom prompts, select the output format, and choose which LLM to run it on.
It even lets you adjust the color scheme and display both the original and translated subtitles side by side.

I originally built it as a tool for language learning-customizable prompts make it easy to guide the translation style.
For instance, you can ask for more literal translations.

### Features:
- **Configurable Automation**: Define your own prompts to process subtitles.
- **Modularity**: Choose your preferred AI model.
- **SRT support**: Easily import & export subtitles to SRT format.

## Project Status

It was built over a few days but is designed to be modular and extendable, allowing for future enhancements and integrations.

You can easily set up and run this project entirely locally.

## Example Results

Below is an example of a before / after running GoSubAI with an example configuration:

![Example Result](images/demo_1.png)

## Prerequisites

Before using GoSubAI, ensure you have the following installed:

- **Golang**: Download and install Golang from [the official website](https://golang.org/dl/).
- **Ollama**: Download and install Ollama from [Ollama's website](https://ollama.com/).

To run Ollama in server mode, use the following command:

```sh
ollama serve
```

## Example Commands

### Generate subtitles

To produce subtitles, it's essential to specify a configuration that outlines the process to be executed on the subtitle files.
Additionally, provide a SubRip (SRT) file containing the subtitles.

Example:

```sh
go run ./cmd/GoSubAI generate -config ./config/translate_to_eng.json -input ./data/HVOR_BLIR_DET_AV_PENGA.srt
```

### Run Unit Tests

To run unit tests, use the following command:

```sh
go clean -testcache; go test ./...
```

## Example configuration

The following configuration will include the original text as well as the translated text:

```json
{
  "Model": "mistral",
  "PropertyName": "translated_text",
  "Prompt": "Translate this to english and keep it short: '{TEXT}'",
  "Template": "{TEXT}\n----\n{GENERATED_TEXT}"
}
```

If you only want to include the translation you can change the `Template` to `"Template": "{GENERATED_TEXT}"`.

For some subtitles, it can be useful to include context from the previous line.
Here’s an example configuration that does this:

```json
{
  "Model": "deepseek-r1:8b",
  "PropertyName": "translated_text",
  "Prompt": "Consider the previous text for context ONLY: '{PREVIOUS_TEXT} {TEXT}'. Translate this to english and keep it short, and don't include the context in your response: '{TEXT}'",
  "Template": "{GENERATED_TEXT}"
}
```