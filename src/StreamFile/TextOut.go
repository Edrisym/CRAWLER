package StreamFile

import (
	"fmt"
	"os"
	"strings"
)

func TextOut(json []byte, drugName string) {

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

	err := os.WriteFile(filePath, json, 0644)

	//err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		fmt.Println("Error saving TEXT:", err)
		return
	}

	fmt.Println("TEXT saved successfully at", filePath)
}
