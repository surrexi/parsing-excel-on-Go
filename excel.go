package main

import (
    "log"
    "github.com/Luxurioust/excelize"
)

// Количество строк листа
var numberRowsSheet = 0

// Функция чтения файла
// path string - путь до файла
func ReadFile(path string)  {
    xlFile, err := excelize.OpenFile(path)
    CheckError(err)
    log.Println("Excel file: ", path)
    log.Println(" - start updates")
    ReadSheet(xlFile.GetRows("sheet1"))
}

// Чтение листа
func ReadSheet(sheetData [][]string)  {
    numberRowsSheet = len(sheetData)
    i := 2
    for i < numberRowsSheet {
        link := GetLinkFromUrl(sheetData[i][1])
        data := GetMetaSeo(link)
        UpdateSeoInfo(data)
        i++
    }
    log.Print(" - success")
}
