package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Written by Aimen Hamed

func main() {
	// If args aren't provided exit with error message
	if len(os.Args) < 3 {
		log.Println("usage: go run txt2go [filename] [P or L]")
		os.Exit(2)
	}

	// extract file name from args and prefix for file
	file := os.Args[1]
	fileprefix := strings.TrimSuffix(file, ".txt")

	// extract orientation for format of pdf
	orientation := os.Args[2]

	// read the txt file in bytes
	content, err := ioutil.ReadFile(file)

	// if the file cannot be found print an error message
	if err != nil {
		log.Fatalf("%s file not found.", file)
	}

	// forms new pdf in A4 size with text in size 12 arial
	pdf := gofpdf.New(orientation, "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	// output file into the text file name but as pdf
	pdfFileName := fileprefix + ".pdf"
	pdf.OutputFileAndClose(pdfFileName)

	// success message
	fmt.Println("PDF Created: ", pdfFileName)
}
