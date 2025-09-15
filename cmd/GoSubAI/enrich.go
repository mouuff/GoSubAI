package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/mouuff/GoSubAI/pkg/brain"
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

	config   string
	datafile string
}

// Name gets the name of the command
func (cmd *GenerateCmd) Name() string {
	return "generate"
}

// Init initializes the command
func (cmd *GenerateCmd) Init(args []string) error {
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.StringVar(&cmd.config, "config", "", "configuration file (required)")
	cmd.flagSet.StringVar(&cmd.datafile, "datafile", "", "file used to load and store data (required)")
	return cmd.flagSet.Parse(args)
}

// Run runs the command
func (cmd *GenerateCmd) Run() error {
	log.Println("Enriching subtitles...")

	if cmd.config == "" {
		log.Println("Please pass the configuration file using -config")
		log.Println("Here is an example configuration:")
		return errors.New("-config parameter required")
	}
	if cmd.datafile == "" {
		log.Println("Please specify a data file using -datafile (e.g. -datafile data.json)")
		return errors.New("-datafile parameter required")
	}

	brain, err := brain.NewOllamaBrain("mistral")

	if err != nil {
		return fmt.Errorf("failed to create brain: %v", err)
	}

	fmt.Println("Loading configuration from", cmd.config)
	fmt.Println("Generating", cmd.datafile)

	r, err := brain.GenerateString(context.Background(), "translated_text", "Translate this to norwegian: This is an example prompt")

	if err != nil {
		return fmt.Errorf("failed to generate: %v", err)
	}

	fmt.Println(r)
	return nil
}
