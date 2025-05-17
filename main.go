package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/p1xelse/merge-pdf-cli/internal"
)

type config struct {
	FilePaths []string
	OutPath   string
}

var cfg config

func init() {
	flag.Func("file", "One of the files to merge", func(s string) error {
		cfg.FilePaths = append(cfg.FilePaths, s)
		return nil
	})

	flag.StringVar(&cfg.OutPath, "out", "merge.pdf", "Output file path (/path/to/file)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintln(os.Stderr, "pdf-merge -file=file1.pdf -file=file2.pdf -out=/path/to/merged.pdf")
	}

	flag.Parse()
}

func main() {
	if err := validateCfg(); err != nil {
		fmt.Fprintf(os.Stderr, "Validation error: %s\n", err.Error())
		return
	}

	fmt.Println("Starting merging process...")

	if err := internal.MergePdf(cfg.FilePaths, cfg.OutPath); err != nil {
		fmt.Fprintf(os.Stderr, "Merge error: %s\n", err.Error())
		return
	}

	fmt.Println("Merge successful.")
}

func validateCfg() error {
	if len(cfg.FilePaths) == 0 {
		return errors.New("no files specified")
	}

	for _, path := range cfg.FilePaths {
		if !strings.Contains(path, ".pdf") {
			return errors.New("file must be a PDF")
		}
	}

	if cfg.OutPath == "" {
		return errors.New("no output path specified")
	}

	return nil
}
