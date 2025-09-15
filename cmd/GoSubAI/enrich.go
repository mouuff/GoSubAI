package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

// Ms describes the generate-trend subcommand
// This command is used to generate trend
type EnrichCmd struct {
	flagSet *flag.FlagSet

	config   string
	datafile string
}

// Name gets the name of the command
func (cmd *EnrichCmd) Name() string {
	return "enrich"
}

// Init initializes the command
func (cmd *EnrichCmd) Init(args []string) error {
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.StringVar(&cmd.config, "config", "", "configuration file (required)")
	cmd.flagSet.StringVar(&cmd.datafile, "datafile", "", "file used to load and store data (required)")
	return cmd.flagSet.Parse(args)
}

// Run runs the command
func (cmd *EnrichCmd) Run() error {
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

	fmt.Println("Loading configuration from", cmd.config)
	fmt.Println("Enriching", cmd.datafile)
	return nil
}
