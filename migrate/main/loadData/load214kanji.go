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
	"strconv"
	"strings"
)

func main() {
	jsonFile, err := os.Open("./214kanji.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open file data successfully")

	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]map[string][]string

	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println("convert successfully")
	var dataAll []models.Alphabet
	for key, value := range result {
		stroke := getStroke(key)
		dataAll = append(dataAll, getData(value, stroke)...)
	}
	fmt.Println("gen data successfully")
	migrateSetting.Setup()
	migrateModel.DBSetup()
	clearData()
	res := migrateModel.DB.Create(&dataAll)

	if res != nil {
		fmt.Println("Successfully")
	} else {
		fmt.Println("Unsuccessfully")
	}
	defer jsonFile.Close()
}

func getStroke(v string) int {
	data := strings.Split(v, "_")
	result, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	return result
}

var Type = map[string]int{"HIRAGANA": 1, "KATAKANA": 2, "KANJI": 3}

func getData(i map[string][]string, stroke int) []models.Alphabet {
	var result []models.Alphabet
	for key, value := range i {
		var note = ""
		if len(value) >= 3 {
			note = value[2]
		}
		result = append(result, models.Alphabet{Type: Type["KANJI"], Word: key, Stroke: stroke, Meaning: value[0], Detail: value[1], Note: note})
	}
	return result
}

func clearData() {
	query := "DELETE from alphabets "
	migrateModel.DB.Exec(query)
}
