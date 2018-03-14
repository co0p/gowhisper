package gowhisper

import (
	"errors"
	"flag"
	"net/url"
	"os"
)

type Flags struct {
	NotifyURL         string
	PollingInterval   int
	ConfigurationFile string
}

func ParseFlags(args []string) (Flags, error) {
	var flags Flags

	flag.StringVar(&flags.NotifyURL, "notifyURL", "", "URL to the notification service")
	flag.StringVar(&flags.ConfigurationFile, "configurationFile", "", "path/to/configuration file")
	flag.IntVar(&flags.PollingInterval, "pollingInterval", 60, "polling interval in seconds (10 - 360)")
	flag.CommandLine.Parse(args)

	if _, err := url.ParseRequestURI(flags.NotifyURL); err != nil {
		return Flags{}, errors.New("NotifyURL should be a valid URL")
	}

	if _, err := os.Stat(flags.ConfigurationFile); err != nil {
		return Flags{}, errors.New("ConfigurationFile should point to a file")
	}

	if flags.PollingInterval < 10 || flags.PollingInterval > 360 {
		return Flags{}, errors.New("pollingInterval should be between 10 and 360")
	}
	return flags, nil
}
