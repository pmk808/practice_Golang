package main

import (
    "fmt"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "github.com/xuri/excelize/v2"
)

func main() {
    r := gin.Default()
    r.POST("/upload", handleExcelUpload)
    r.Run(":8080")
}

func handleExcelUpload(c *gin.Context) {
    // File upload handling
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }

    // Validate file extension
    if filepath.Ext(file.Filename) != ".xlsx" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Please upload an Excel file (.xlsx)"})
        return
    }

    // Save the uploaded file
    if err := c.SaveUploadedFile(file, file.Filename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    // Open the Excel file
    f, err := excelize.OpenFile(file.Filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open Excel file"})
        return
    }
    defer f.Close()

    // Get all sheet names
    sheets := f.GetSheetList()

    // Process each sheet
    var result []map[string]interface{}
    for _, sheet := range sheets {
        sheetData, err := processSheet(f, sheet)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error processing sheet %s: %v", sheet, err)})
            return
        }
        result = append(result, sheetData)
    }

    c.JSON(http.StatusOK, gin.H{"data": result})
}

func processSheet(f *excelize.File, sheet string) (map[string]interface{}, error) {
    rows, err := f.GetRows(sheet)
    if err != nil {
        return nil, err
    }

    if len(rows) == 0 {
        return map[string]interface{}{"sheet": sheet, "data": []map[string]interface{}{}}, nil
    }

    headers := rows[0]
    var data []map[string]interface{}

    for _, row := range rows[1:] {
        rowData := make(map[string]interface{})
        for i, cell := range row {
            if i < len(headers) {
                rowData[headers[i]] = cell
            }
        }
        data = append(data, rowData)
    }

    return map[string]interface{}{"sheet": sheet, "data": data}, nil
}
