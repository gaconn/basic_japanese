package main

import (
	"encoding/json"
	"fmt"
	"github.com/quan12xz/basic_japanese/migrate/migrateSetting"
	migrateModel "github.com/quan12xz/basic_japanese/migrate/models"
	"github.com/quan12xz/basic_japanese/models"
	"io"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("./hirakata.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open file data successfully")

	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]map[string]string

	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println("convert successfully")
	var dataAll []models.Alphabet
	for key, value := range result {
		dataAll = append(dataAll, getDataHiraKata(value, key)...)
	}
	fmt.Println("gen data successfully")
	migrateSetting.Setup()
	migrateModel.DBSetup()
	clearHiraKataData()
	res := migrateModel.DB.Create(&dataAll)

	if res != nil {
		fmt.Println("Successfully")
	} else {
		fmt.Println("Unsuccessfully")
	}
	defer jsonFile.Close()
}

var typeW = map[string]int{"HIRAGANA": 1, "HIRAGANA_COMBINE": 11, "KATAKANA": 2, "KATAKANA_COMBINE": 22, "KANJI": 3}

func getDataHiraKata(i map[string]string, key string) []models.Alphabet {
	var result []models.Alphabet
	var typeWord int
	if key == "hiragana" {
		typeWord = typeW["HIRAGANA"]
	} else if key == "hiragana_combine" {
		typeWord = typeW["HIRAGANA_COMBINE"]
	} else if key == "katakana" {
		typeWord = typeW["KATAKANA"]
	} else if key == "katakana_combine" {
		typeWord = typeW["KATAKANA_COMBINE"]
	}
	for key, value := range i {
		result = append(result, models.Alphabet{Type: typeWord, Word: value, Meaning: key})
	}
	return result
}

func clearHiraKataData() {
	query := "DELETE from alphabets WHERE type in (1,11,2,22)"
	migrateModel.DB.Exec(query)
}
