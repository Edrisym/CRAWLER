package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func excelized() {

	f := excelize.NewFile()

	sheetName := "Products"
	index, _ := f.NewSheet(sheetName)

	// Set headers
	headers := []string{
		"Persian Name",
		"English Name",
		"Brand Owner",
		"License Holder",
		"Price",
		"Packaging",
		"Product Code",
		"Generic Code",
	}

	// Add headers to the first row
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i))
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			log.Fatal(err)
		}
	}

	// Set the active sheet
	f.SetActiveSheet(index)

	// Save the file
	if err := f.SaveAs("ProductTemplate.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Excel template created successfully.")
}
