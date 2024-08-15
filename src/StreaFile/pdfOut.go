package StreaFile

import (
	"fmt"
	"github.com/jung-kurt/gofpdf/v2"
	"os"
	"strings"
)

func PdfOut(json []byte, drugName string) {

	folderPath := "output"

	words := []string{"drugs", "-", drugName, ".txt"}
	fileName := strings.TrimSpace(strings.Join(words, ""))

	filePath := fmt.Sprintf("%s/%s", folderPath, fileName)

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	}
	text := string(json)

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, text)

	err := os.WriteFile(filePath, json, 0644)

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		fmt.Println("Error saving TEXT:", err)
		return
	}

	fmt.Println("TEXT saved successfully at", filePath)
}
