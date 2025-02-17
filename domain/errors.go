package domain

import "errors"

var (
	ErrRequireAtLeastTwoFile   = errors.New("at least two files are required to merge")
	ErrCannotMerge             = errors.New("cannot merge files")
	ErrCannotSplit             = errors.New("cannot split file")
	ErrCannotCompress          = errors.New("cannot compress file")
	ErrInvalidCompressionLevel = errors.New("invalid compression level")
)
