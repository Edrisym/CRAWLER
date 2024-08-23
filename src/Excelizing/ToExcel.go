package Excelizing

import (
	"CrawlerBot/Product"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

func ToExcel(jsonData []byte, drugName string) {

	// Unmarshal the JSON data into a slice of Product structs
	var products []Product.Product
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
		"Licence Date",
	}

	// Add headers to the first row
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i))
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			log.Fatal(err)
		}
	}

	for i, product := range products {
		row := i + 2

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), product.PersianName)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), product.EnglishName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), product.BrandOwner)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), product.LicenseHolder)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), product.Price)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), product.Packaging)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), product.ProductCode)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), product.GenericCode)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), product.ProductDetails.LicenceDate)
	}

	f.SetActiveSheet(index)

	words := []string{"Excel-Drug-", drugName, ".xlsx"}
	fileName := strings.TrimSpace(strings.Join(words, ""))

	if err := f.SaveAs(fileName); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Excel file created successfully.")
}
