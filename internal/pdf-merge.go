package internal

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func MergePdf(paths []string, outPath string) error {
	config := model.NewDefaultConfiguration()

	err := api.MergeCreateFile(paths, outPath, true, config)
	if err != nil {
		return fmt.Errorf("error merging PDF files: %v", err)
	}

	return nil
}
