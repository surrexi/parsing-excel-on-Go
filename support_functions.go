package main

import (
	"log"
	"strings"
)

// Проверка на наличие ошибок
func CheckError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func GetLinkFromUrl(str string) string {
	return strings.Replace(str, "https://stalbeton.pro/catalog-", "", -1)
}