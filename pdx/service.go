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

/*
Merge function merges multiple PDF files into a single PDF file.
The inputFiles parameter is a slice of strings that contains the paths to the PDF files to merge.
The outputFile parameter is the path to the output PDF file.
*/
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

/*
Split function splits the PDF file into multiple files with the specified number of pages.
The pageNrs parameter specifies the number of pages to split the PDF file into.
*/
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

/*
Compress function compresses the PDF file with the specified level.
The level can be one of the following(default is domain.RecommendedCompressionLevel):

- domain.LowCompression (Low optimization)

- domain.RecommendedCompressionLevel (medium optimization)

- domain.ExtremelyHighCompression (high optimization)
*/
func (s *PDXService) Compress(inputFile string, outputFile string, level string) error {
	err := s.pdf.Compress(inputFile, outputFile, domain.ExtremelyHighCompression)
	if err != nil {
		log.Println(err)
		return domain.ErrCannotCompress
	}
	return nil
}
