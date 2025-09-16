package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/mouuff/GoSubAI/internal"
	"github.com/mouuff/GoSubAI/pkg/brain"
	"github.com/mouuff/GoSubAI/pkg/generator"
	"github.com/mouuff/GoSubAI/pkg/parser"
	"github.com/mouuff/GoSubAI/pkg/writer"
)

type GeneratorConfig struct {
	Model        string
	PropertyName string
	Prompt       string
	Template     string
	Debug        bool
}

// Ms describes the generate-trend subcommand
// This command is used to generate trend
type GenerateCmd struct {
	flagSet *flag.FlagSet

	config string
	input  string
	output string
}

// Name gets the name of the command
func (cmd *GenerateCmd) Name() string {
	return "generate"
}

// Init initializes the command
func (cmd *GenerateCmd) Init(args []string) error {
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.StringVar(&cmd.config, "config", "", "configuration file (required)")
	cmd.flagSet.StringVar(&cmd.input, "input", "", "file used to load SRT data (required)")
	cmd.flagSet.StringVar(&cmd.output, "output", "", "file used to write SRT data")
	return cmd.flagSet.Parse(args)
}

// Run runs the command
func (cmd *GenerateCmd) Run() error {
	if cmd.config == "" {
		log.Println("Please pass the configuration file using -config")
		log.Println("Here is an example configuration:")
		printConfigurationTemplate()
		return errors.New("-config parameter required")
	}
	if cmd.input == "" {
		log.Println("Please specify a data file using -input (e.g. -input data.srt)")
		return errors.New("-input parameter required")
	}

	if cmd.output == "" {
		cmd.output = internal.AddPrefixToFilename(cmd.input, "_generated")
		log.Printf("No '-output' file specified, using default: %s\n", cmd.output)
	}

	var config GeneratorConfig
	err := internal.ReadFromJson(cmd.config, &config)
	if err != nil {
		return err
	}

	parser := &parser.SrtParser{}
	writer := &writer.SrtWriter{}

	subtitleData, err := parser.Parse(cmd.input)
	if err != nil {
		return fmt.Errorf("failed to read subtitle data: %v", err)
	}

	brain, err := brain.NewOllamaBrain("mistral")

	if err != nil {
		return fmt.Errorf("failed to create brain: %v", err)
	}

	generator := &generator.SubtitleGenerator{
		Context:       context.Background(),
		Brain:         brain,
		PropertyName:  config.PropertyName,
		Prompt:        config.Prompt,
		Template:      config.Template,
		Debug:         config.Debug,
		SubstitleData: subtitleData,
	}

	result, err := generator.Generate()
	if err != nil {
		return err
	}

	return writer.Write(cmd.output, result)
}

func printConfigurationTemplate() {
	configTemplate := &GeneratorConfig{
		Model:        "mistral",
		PropertyName: "translated_text",
		Prompt:       "Translate this to english: '{TEXT}'",
		Template:     "{TEXT}\n{GENERATED_TEXT}",
		Debug:        false,
	}

	jsonData, err := json.MarshalIndent(configTemplate, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
