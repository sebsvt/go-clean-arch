package pdx

import (
	"log"

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

func (s *PDXService) Merge(inputFiles []string, outputFile string) error {
	if len(inputFiles) < 2 {
		return domain.ErrRequireAtLeastTwoFile
	}
	err := s.pdf.Merge(inputFiles, outputFile)
	if err != nil {
		log.Println(err)
		return domain.ErrCannotMerge
	}
	return nil
}

func (s *PDXService) Split(inputFile string, outputDir string, pageNrs int) error {
	if pageNrs < 1 {
		return domain.ErrRequireAtLeastOnePage
	}
	err := s.pdf.Split(inputFile, outputDir, pageNrs)
	if err != nil {
		log.Println(err)
		return domain.ErrCannotSplit
	}
	return nil
}

func (s *PDXService) Compress(inputFile string, outputFile string, level string) error {
	err := s.pdf.Compress(inputFile, outputFile, level)
	if err != nil {
		log.Println(err)
		return domain.ErrCannotCompress
	}
	return nil
}
