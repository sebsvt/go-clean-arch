package pdx_test

import (
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/adapter"
	"github.com/bxcodec/go-clean-arch/pdx"
)

func TestMergeFunction(t *testing.T) {
	type testCase struct {
		name        string
		inputFiles  []string
		outputFile  string
		expectedErr error
	}
	testcases := []testCase{
		{
			name:        "Merging two files into one file should not return an error",
			inputFiles:  []string{"../assets/cv.pdf", "../assets/file2.pdf"},
			outputFile:  "../assets/output.pdf",
			expectedErr: nil,
		},
		{
			name:        "Merging one file into one file should return an error",
			inputFiles:  []string{"../assets/cv.pdf"},
			outputFile:  "../assets/output.pdf",
			expectedErr: domain.ErrRequireAtLeastTwoFile,
		},
		{
			name:        "Merging two files into one file with invalid output file should return an error",
			inputFiles:  []string{"../assets/cv.pdf", "../assets/file2.pdf"},
			outputFile:  "",
			expectedErr: domain.ErrCannotMerge,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			pdf := pdx.New(adapter.NewPDFCPUAdapter())
			err := pdf.Merge(tc.inputFiles, tc.outputFile)
			if err != tc.expectedErr {
				t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type testCase struct {
		name        string
		inputFiles  string
		outputDir   string
		page        int
		expectedErr error
	}
	testcases := []testCase{
		{
			name:        "Splitting one file into two files should not return an error",
			inputFiles:  "../assets/output.pdf",
			outputDir:   "../assets/out",
			page:        1,
			expectedErr: nil,
		},
		{
			name:        "Splitting one file into two files with page 0 should return an error",
			inputFiles:  "../assets/output.pdf",
			outputDir:   "../assets/out",
			page:        0,
			expectedErr: domain.ErrRequireAtLeastOnePage,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			pdf := pdx.New(adapter.NewPDFCPUAdapter())
			err := pdf.Split(tc.inputFiles, tc.outputDir, tc.page)
			if err != tc.expectedErr {
				t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestCompress(t *testing.T) {
	type testCase struct {
		name        string
		inputFiles  string
		outputFile  string
		level       string
		expectedErr error
	}
	testcases := []testCase{
		{
			name:        "Compressing one file into one file should not return an error",
			inputFiles:  "../assets/output.pdf",
			outputFile:  "../assets/output_compressed.pdf",
			level:       domain.RecommendedCompressionLevel,
			expectedErr: nil,
		},
		{
			name:        "Compressing invalid file should return an error",
			inputFiles:  "../assets/output.pdfx",
			outputFile:  "../assets/output_compressed.pdf",
			level:       domain.RecommendedCompressionLevel,
			expectedErr: domain.ErrCannotCompress,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			pdf := pdx.New(adapter.NewPDFCPUAdapter())
			err := pdf.Compress(tc.inputFiles, tc.outputFile, tc.level)
			if err != tc.expectedErr {
				t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
