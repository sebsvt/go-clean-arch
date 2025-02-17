package main

import (
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/adapter"
	"github.com/bxcodec/go-clean-arch/pdx"
)

func main() {
	// Initialize the PDF Service
	pdx_srv := pdx.New(adapter.NewPDFCPUAdapter())
	pdx_srv.Merge([]string{"file1.pdf", "file2.pdf"}, "output.pdf")
	pdx_srv.Split("input.pdf", "output", 10)
	pdx_srv.Compress("input.pdf", "output.pdf", domain.LowCompression)
}
