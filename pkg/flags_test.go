package gowhisper_test

import (
	"flag"
	"os"
	"strconv"
	"testing"

	"github.com/co0p/gowhisper/pkg"
)

func ResetForTesting() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func Test_ParseFlags_Should_ErrOnMissingFlagsGiven(t *testing.T) {
	ResetForTesting()
	args := []string{""}

	_, err := gowhisper.ParseFlags(args)
	if err == nil {
		t.Errorf("expected err not to be nil, got %v", err)
	}
}

func Test_ParseFlags_Should_ErrOnInvalidPollingIntervalGiven(t *testing.T) {
	ResetForTesting()

	args := []string{"-pollingInterval", "0", "-configurationFile", "/tmp"}
	_, err := gowhisper.ParseFlags(args)
	if err == nil {
		t.Errorf("expected err not to be nil, got '%v'", err)
	}
}

func Test_ParseFlags_Should_ErrOnInvalidPortGiven(t *testing.T) {
	ResetForTesting()

	args := []string{"-pollingInterval", "100", "-configurationFile", "/tmp", "-port", "0"}
	_, err := gowhisper.ParseFlags(args)
	if err == nil {
		t.Errorf("expected err not to be nil, got '%v'", err)
	}
}

func Test_ParseFlags_Should_ErrOnInvalidConfigurationFileGiven(t *testing.T) {
	ResetForTesting()

	args := []string{"-pollingInterval", "60", "-configurationFile", " does not exist"}
	_, err := gowhisper.ParseFlags(args)
	if err == nil {
		t.Errorf("expected err not to be nil, got '%v'", err)
	}
}

func Test_ParseFlags_Should_ReturnParsedFlags(t *testing.T) {
	ResetForTesting()
	pollingInterval := "10"
	configurationFile := "../clients.json"
	port := "7777"

	args := []string{"-pollingInterval", pollingInterval, "-configurationFile", configurationFile, "-port", port}
	flags, err := gowhisper.ParseFlags(args)
	if err != nil {
		t.Errorf("expected err to be nil, got '%v'", err)
	}

	if flags.ConfigurationFile != configurationFile {
		t.Errorf("expected ConfigurationFile to be %v, got '%v'", configurationFile, flags.ConfigurationFile)
	}

	pollingIntervalInt, _ := strconv.Atoi(pollingInterval)
	if flags.PollingInterval != pollingIntervalInt {
		t.Errorf("expected PollingInterval to be %v, got '%v'", pollingIntervalInt, flags.PollingInterval)
	}

	portInt, _ := strconv.Atoi(port)
	if flags.Port != portInt {
		t.Errorf("expected Port to be %v, got '%v'", portInt, flags.Port)
	}
}
