package adapter

import (
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type PDFCPUAdapter struct {
	config *model.Configuration
}

func NewPDFCPUAdapter() *PDFCPUAdapter {
	cfg := model.NewDefaultConfiguration()
	return &PDFCPUAdapter{
		config: cfg,
	}
}

func (p *PDFCPUAdapter) Merge(inputFiles []string, outputFile string) error {
	return api.MergeCreateFile(inputFiles, outputFile, false, p.config)
}

func (p *PDFCPUAdapter) Split(inputFile string, outputDir string, pageNrs int) error {
	err := api.SplitFile(inputFile, outputDir, pageNrs, p.config)
	return err
}

func (p *PDFCPUAdapter) Compress(inputFile string, outputFile string, level string) error {
	conf := model.NewDefaultConfiguration()
	conf.Optimize = true
	switch level {
	case domain.ExtremelyHighCompression:
		conf.OptimizeResourceDicts = true
		conf.OptimizeDuplicateContentStreams = true
	case domain.RecommendedCompressionLevel:
		conf.OptimizeResourceDicts = true
		conf.OptimizeDuplicateContentStreams = false
	case domain.LowCompression:
		conf.OptimizeResourceDicts = false
		conf.OptimizeDuplicateContentStreams = false
	default:
		// if level is not recognized, use recommended
		conf.OptimizeResourceDicts = true
		conf.OptimizeDuplicateContentStreams = false

	}
	err := api.OptimizeFile(inputFile, outputFile, conf)
	return err
}
