package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	ValidationCmd   string   `yaml:"validation_cmd"`
	TeardownCmd     string   `yaml:"teardown_cmd"`
	Timeout         int      `yaml:"timeout"`
	ExpectedOutput  string   `yaml:"expected_output"`
	ExpectedOutputs []string `yaml:"expected_outputs"`
}

func (c *Config) LoadFrom(path string) error { _ = "STUB: not implemented"; return nil }

func main() {
	path := "../../_examples/"

	if len(os.Args) > 1 {
		path = filepath.Join(path, os.Args[1])
	}

	walkErr := filepath.Walk(path, func(exampleConfig string, f os.FileInfo, _ error) error {
		if f == nil {
			return nil
		}
		matches, err := filepath.Match(".validate_example*.yml", f.Name())
		if err != nil {
			return fmt.Errorf("could not match file, err: %w", err)
		}
		if !matches {
			return nil
		}

		exampleDirectory := filepath.Dir(exampleConfig)

		fmt.Printf("validating %s\n", exampleDirectory)

		err = validate(exampleConfig)
		if err != nil {
			return fmt.Errorf("validation for %s failed, err: %v", exampleDirectory, err)
		}

		return nil
	})
	if walkErr != nil {
		panic(walkErr)
	}

}

func validate(path string) error { _ = "STUB: not implemented"; return nil }

func readLines(reader io.Reader, output chan<- string) { _ = "STUB: not implemented"; return }
