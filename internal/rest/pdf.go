package rest

// @Summary Compress a PDF file
// @Description Compress a given PDF file with a specified compression level
// @Tags PDF
// @Accept json
// @Produce json
// @Param inputFile body string true "Path to input PDF file"
// @Param outputFile body string true "Path to output compressed PDF file"
// @Param level body string true "Compression level (low, medium, high)"
// @Success 200 {object} map[string]string "Compression successful"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Compression failed"
// @Router /api/v1/compress [post]
func CompressHandler() {}

// @Summary Merge multiple PDF files
// @Description Merge multiple PDF files into a single file
// @Tags PDF
// @Accept json
// @Produce json
// @Param inputFiles body []string true "List of input PDF file paths"
// @Param outputFile body string true "Path to output merged PDF file"
// @Success 200 {object} map[string]string "Merge successful"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Merge failed"
// @Router /api/v1/merge [post]
func MergeHandler() {}

// @Summary Split a PDF file
// @Description Split a PDF file into multiple parts based on page numbers
// @Tags PDF
// @Accept json
// @Produce json
// @Param inputFile body string true "Path to input PDF file"
// @Param outputDir body string true "Path to directory for output split files"
// @Param pageNrs body int true "Number of pages per split file"
// @Success 200 {object} map[string]string "Split successful"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Split failed"
// @Router /api/v1/split [post]
func SplitHandler() {}
