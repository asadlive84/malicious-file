package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run pdf_analysis.go <input.pdf>")
		return
	}

	inputPath := os.Args[1]

	// Open the PDF file.
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Error opening PDF file: %v", err)
	}
	defer file.Close()

	// Create a default configuration.
	config := model.NewDefaultConfiguration()

	// Read the PDF content.
	ctx, err := pdfcpu.Read(file, config)
	if err != nil {
		log.Fatalf("Error reading PDF: %v", err)
	}

	// Iterate over pages and extract text.
	if err := ctx.EnsurePageCount(); err != nil {
		log.Fatalf("Error ensuring page count: %v", err)
	}

	for pageNum := 1; pageNum <= ctx.PageCount; pageNum++ {
		// Extract content stream from the page.
		content, err := pdfcpu.ExtractPageContent(ctx, pageNum)
		if err != nil {
			log.Fatalf("Error extracting content from page %d: %v", pageNum, err)
		}

		// ReadAll reads from the reader until an error or EOF and returns the data as a byte slice.
		_, err = io.ReadAll(content)
		if err != nil {
			log.Fatalf("Error reading content from page %d: %v", pageNum, err)
		}

		// Convert byte slice to string.
		// text := string(data)

		// if err := api.ExportFormFile(text, "test.json", config); err != nil {
		// 	log.Fatalf("Error reading content from page %d: %v", pageNum, err)
		// }

		// fmt.Printf("Text from page %d:\n%s\n", pageNum, text)
	}

	// Access the XRefTable.
	xRefTable := ctx.XRefTable
	if xRefTable == nil {
		log.Fatal("XRefTable is nil")
	}

	// Print information about each object.
	for objectNumber, entry := range xRefTable.Table {
		fmt.Printf("Object %d: %v\n", objectNumber, entry)
	}
}

// Interpret content stream to extract text.
// func interpretContent(content io.Reader) (string, error) {
// 	parser := interpreter.NewParser(content)
// 	processor := interpreter.NewProcessor()

// 	err := processor.Process(parser)
// 	if err != nil {
// 		return "", err
// 	}

// 	return processor.Text(), nil
// }
