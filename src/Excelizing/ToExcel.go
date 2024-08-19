package Excelizing

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

type Product struct {
	PersianName   string `json:"persian_name"`
	EnglishName   string `json:"english_name"`
	BrandOwner    string `json:"brand_owner"`
	LicenseHolder string `json:"license_holder"`
	Price         string `json:"price"`
	Packaging     string `json:"packaging"`
	ProductCode   string `json:"product_code"`
	GenericCode   string `json:"generic_code"`
}

func ToExcel(jsonData []byte, drugName string) {

	// Unmarshal the JSON data into a slice of Product structs
	var products []Product
	if err := json.Unmarshal(jsonData, &products); err != nil {
		log.Fatal(err)
	}

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

	// Populate the sheet with the data
	for i, product := range products {
		row := i + 2 // Start from the second row

		// Set each cell's value
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), product.PersianName)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), product.EnglishName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), product.BrandOwner)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), product.LicenseHolder)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), product.Price)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), product.Packaging)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), product.ProductCode)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), product.GenericCode)
	}

	// Set the active sheet
	f.SetActiveSheet(index)

	words := []string{"Excel-Drug", drugName, ".xlsx"}
	fileName := strings.TrimSpace(strings.Join(words, ""))

	if err := f.SaveAs(fileName); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Excel file created successfully.")
}
