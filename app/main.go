package main

import (
	"log"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/adapter"
	"github.com/bxcodec/go-clean-arch/pdx"
)

func main() {
	// Initialize the PDF Service
	pdx_srv := pdx.New(adapter.NewPDFCPUAdapter())
	err := pdx_srv.Merge([]string{"C:/Users/vithc/Documents/_workspace/go-clean-arch/assets/cv.pdf", "C:/Users/vithc/Documents/_workspace/go-clean-arch/assets/file2.pdf"}, "C:/Users/vithc/Documents/_workspace/go-clean-arch/assets/output.pdf")
	if err != nil {
		log.Println(err)
	}
	err = pdx_srv.Split("C:/Users/vithc/Documents/_workspace/go-clean-arch/assets/output.pdf", "C:/Users/vithc/Documents/_workspace/go-clean-arch/assets/out", 1)
	if err != nil {
		log.Println(err)
	}
	err = pdx_srv.Compress("C:/Users/vithc/Documents/_workspace/go-clean-arch/assets/output.pdf", "C:/Users/vithc/Documents/_workspace/go-clean-arch/assets/output_compressed.pdf", domain.ExtremelyHighCompression)
	if err != nil {
		log.Println(err)
	}
}
