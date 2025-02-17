package pdx

import "github.com/bxcodec/go-clean-arch/domain"

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
		return domain.ErrCannotMerge
	}
	return nil
}

func (s *PDXService) Split(inputFile string, outputDir string, pageNrs int) error {
	return s.pdf.Split(inputFile, outputDir, pageNrs)
}

func (s *PDXService) Compress(inputFile string, outputFile string, level string) error {
	return s.pdf.Compress(inputFile, outputFile, level)
}
