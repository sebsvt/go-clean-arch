package service

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
	return s.pdf.Merge(inputFiles, outputFile)
}

func (s *PDXService) Split(inputFile string, outputDir string, pageNrs int) error {
	return s.pdf.Split(inputFile, outputDir, pageNrs)
}

func (s *PDXService) Compress(inputFile string, outputFile string, level string) error {
	return s.pdf.Compress(inputFile, outputFile, level)
}
