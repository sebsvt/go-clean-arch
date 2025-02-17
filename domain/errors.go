package domain

import "errors"

var (
	ErrRequireAtLeastTwoFile = errors.New("at least two files are required to merge")
	ErrCannotMerge           = errors.New("cannot merge files")
	ErrCannotSplit           = errors.New("cannot split file")
	ErrCannotCompress        = errors.New("cannot compress file")
	ErrRequireAtLeastOnePage = errors.New("at least one page is required to split")
	ErrRequireOutputFile     = errors.New("output file is required to compress")
	ErrRequireOutputDir      = errors.New("output directory is required to split")
)
