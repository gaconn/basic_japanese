package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/models"
	"github.com/quan12xz/basic_japanese/utils"
	"strconv"
	"strings"
)

var TypeWord = map[string]int{"HIRAGANA": 1, "HIRAGANA_COMBINE": 11, "KATAKANA": 2, "KATAKANA_COMBINE": 22, "KANJI": 3}

func GetAllByType(r *gin.Context) {
	strType := strings.ToUpper(r.Param("type"))
	intType, ok := TypeWord[strType]
	if !ok {
		utils.SendResponse(r, 400, "Type not exist", nil)
	}
	result, err := models.GetWordByType(intType)
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}
	utils.SendResponse(r, 200, "Successfully", result)
}

func GetByID(r *gin.Context) {
	strID := r.Param("id")
	intID, err := strconv.ParseInt(strID, 10, 64)

	if err != nil {
		utils.SendResponse(r, 400, "ID invalid", nil)
	}
	result, err := models.GetWordByID(intID)

	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}
	utils.SendResponse(r, 200, "Successfully", result)
}

func Update(r *gin.Context) {
	var err error

	var alphabet = models.Alphabet{}
	err = r.Bind(&alphabet)

	if err != nil || alphabet.ID == 0 {
		utils.SendResponse(r, 400, "Data invalid", nil)
	}

	result, err := models.UpdateWord(&alphabet)
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}
	utils.SendResponse(r, 200, "Successfully", result)
}
