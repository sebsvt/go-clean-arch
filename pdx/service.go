package pdx

import (
	"log"
	"path/filepath"

	"github.com/bxcodec/go-clean-arch/domain"
)

type PDFPort interface {
	Merge(inputFiles []string, outputFile string) error
	Split(inputFile string, outputDir string, pageNrs int) error
	Compress(inputFile string, outputFile string, level string) error
}

type PDXService struct {
	pdf PDFPort
}

func New(pdf PDFPort) *PDXService {
	return &PDXService{pdf: pdf}
}

func resolvePath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Failed to resolve path %s: %v", path, err)
	}
	return absPath
}

/*
Merge merges multiple PDF files into a single PDF.
*/
func (s *PDXService) Merge(inputFiles []string, outputFile string) error {
	if len(inputFiles) < 2 {
		return domain.ErrRequireAtLeastTwoFile
	}

	absInputFiles := make([]string, len(inputFiles))
	for i, file := range inputFiles {
		absInputFiles[i] = resolvePath(file)
	}
	absOutputFile := resolvePath(outputFile)

	err := s.pdf.Merge(absInputFiles, absOutputFile)
	if err != nil {
		log.Println(err)
		return domain.ErrCannotMerge
	}
	return nil
}

/*
Split splits a PDF file into multiple files.
*/
func (s *PDXService) Split(inputFile string, outputDir string, pageNrs int) error {
	if pageNrs < 1 {
		return domain.ErrRequireAtLeastOnePage
	}

	absInputFile := resolvePath(inputFile)
	absOutputDir := resolvePath(outputDir)

	err := s.pdf.Split(absInputFile, absOutputDir, pageNrs)
	if err != nil {
		log.Println(err)
		return domain.ErrCannotSplit
	}
	return nil
}

/*
Compress compresses a PDF file with a specified level.
*/
func (s *PDXService) Compress(inputFile string, outputFile string, level string) error {
	// Auto-resolve paths
	absInputFile := resolvePath(inputFile)
	absOutputFile := resolvePath(outputFile)

	err := s.pdf.Compress(absInputFile, absOutputFile, level)
	if err != nil {
		log.Println(err)
		return domain.ErrCannotCompress
	}
	return nil
}
