package gowhisper

import (
	"errors"
	"flag"
	"os"
)

// Flags contains the passed in settings via flags
type Flags struct {
	NotifyURL         string
	PollingInterval   int
	ConfigurationFile string
	Port              int
}

func ParseFlags(args []string) (Flags, error) {
	var flags Flags

	flag.StringVar(&flags.ConfigurationFile, "configurationFile", "", "path/to/configuration file")
	flag.IntVar(&flags.PollingInterval, "pollingInterval", 60, "polling interval in seconds (10 - 360)")
	flag.IntVar(&flags.Port, "port", 8080, "port to serve status page on")
	flag.CommandLine.Parse(args)

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	if _, err := os.Stat(flags.ConfigurationFile); err != nil {
		return Flags{}, errors.New("-configurationFile should point to a file")
	}

	if flags.PollingInterval < 10 || flags.PollingInterval > 360 {
		return Flags{}, errors.New("-pollingInterval should be between 10 and 360")
	}

	if flags.Port < 80 || flags.Port > 65555 {
		return Flags{}, errors.New("-port should be between 80 and 65555")
	}

	return flags, nil
}
