package StreaFile

import (
	"fmt"
	"github.com/jung-kurt/gofpdf/v2"
	"os"
)

func PdfOut(json []byte) {
	// Define the folder where you want to save the PDF
	folderPath := "output"
	fileName := "drugs.pdf"
	filePath := fmt.Sprintf("%s/%s", folderPath, fileName)

	// Create the folder if it doesn't exist
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, string(json))

	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("PDF saved successfully at", filePath)
}
