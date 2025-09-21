
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
  "HostUrl": "default",
  "Model": "mistral",
  "PropertyName": "translated_text",
  "SystemPrompt": "You are a subtitle translation assistant. Your only task is to translate subtitles into the target language specified by the user. Subtitles may contain incomplete sentences-when that happens, translate them literally without trying to complete or alter their meaning. Always keep the translation faithful to the original text and do not add explanations or extra words.",
  "Prompt": "Translate this to english: '{TEXT}'",
  "Template": "{TEXT}\n----\n{GENERATED_TEXT}",
  "Debug": true
}
```

If you only want to include the translation you can change the `Template` to `"Template": "{GENERATED_TEXT}"`.

For some subtitles, it can be useful to include context from the previous line.
Here’s an example configuration that does this:

```json
{
  "HostUrl": "default",
  "Model": "llama3.2",
  "PropertyName": "translated_text",
  "SystemPrompt": "You are a subtitle translation assistant. Translate subtitles into the target language specified by the user.\n\n- Translate the text literally.\n- Do not add, remove, or guess words.\n- If the text is incomplete, translate it as-is.\n\n⚠️ Very important: The tag ##TAG## must always be copied exactly, in the same position.\n- Never remove it.\n- Never move it.\n- Never translate it.\n- Output is invalid if ##TAG## is missing or changed.\n\n### Examples\n\nInput: \"Bonjour##TAG##comment tu vas aujourd'hui?\"\nOutput: \"Hello##TAG##how are you today?\"\n\nInput: \"##TAG##Oui, j'arrive.\"\nOutput: \"##TAG##Yes, I'm coming.\"\n\nInput: \"Non##TAG##pas du tout.\"\nOutput: \"No##TAG##not at all.\"",
  "Prompt": "Translate this to english: '{PREVIOUS_TEXT} ##TAG## {TEXT}'",
  "Template": "{TEXT}\n----\n{GENERATED_TEXT}",
  "Regex": "##TAG##(.*)",
  "RegexRetryLimit": 25,
  "Debug": true
}
```