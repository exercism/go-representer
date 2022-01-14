package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/exercism/go-representer/representer"
	"github.com/namsral/flag"
)

type options struct {
	exercise     string
	solutionPath string
	outputPath   string
}

func getOptions() (options, bool) {
	var opts options

	flag.StringVar(&opts.exercise, "exercise", "", "exercise slug (e.g. 'two-fer')")
	flag.StringVar(&opts.solutionPath, "solution", "", "path to solution to be processed")
	flag.StringVar(&opts.outputPath, "output", "./", "path for output files")

	flag.Parse()

	if opts.exercise == "" || opts.solutionPath == "" {
		args := flag.Args()
		if len(args) < 2 {
			flag.Usage()
			return opts, false
		}
		opts.exercise, opts.solutionPath = args[0], args[1]
		if 2 < len(args) && opts.outputPath == "./" {
			opts.outputPath = args[2]
		}
	}

	return opts, true
}

func main() {
	opts, ok := getOptions()
	if !ok {
		os.Exit(3)
	}

	err := run(opts)
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
}

func run(opts options) error {
	log.Printf("Creating representation for `%s` exercise in folder `%s`\n", opts.exercise, opts.solutionPath)

	repr, err := representer.Extract(opts.solutionPath)
	if err != nil {
		return fmt.Errorf("failed to extract representation: %w", err)
	}

	reprBts, err := repr.RepresentationBytes()
	if err != nil {
		return fmt.Errorf("failed to serialize representation: %w", err)
	}
	mappingBts, err := repr.MappingBytes()
	if err != nil {
		return fmt.Errorf("failed to serialize representation: %w", err)
	}

	return writeOutput(opts, reprBts, mappingBts)
}

func writeOutput(opts options, representation []byte, mapping []byte) error {
	representationFile := path.Join(opts.outputPath, "representation.txt")
	if err := writeFile(representationFile, representation); err != nil {
		return fmt.Errorf("%s: %w", representationFile, err)
	}

	mappingFile := path.Join(opts.outputPath, "mapping.json")
	if err := writeFile(mappingFile, mapping); err != nil {
		return fmt.Errorf("%s: %w", representationFile, err)
	}
	return nil
}

func writeFile(filePath string, bytes []byte) error {
	if err := ioutil.WriteFile(filePath, bytes, 0644); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	log.Printf("Written to %s", filePath)
	return nil
}
