package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/models"
	"github.com/quan12xz/basic_japanese/utils"
)

func GetAllKatakana(r *gin.Context) {
	result, err := models.GetAllKatakana()
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
	}
	utils.SendResponse(r, 200, "Successfully", result)
}

func GetByIDKatakana(r *gin.Context) {
	strID := r.Param("id")
	intID, err := strconv.ParseInt(strID, 10, 64)

	if err != nil {
		utils.SendResponse(r, 400, "ID invalid", nil)
	}
	result, err := models.GetByIDKatakana(intID)

	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
	}
	utils.SendResponse(r, 200, "Successfully", result)
}

func AddKatakana(r *gin.Context) {
	var alphabet = &models.Alphabet{}
	err := r.Bind(alphabet)
	if err != nil {
		log.Fatal(err)
	}
	data, err := models.CreateKatakana(alphabet)
	if err != nil {
		utils.SendResponse(r, 304, err.Error(), nil)
	}
	utils.SendResponse(r, 201, "Add new katakana work successfully", data)
}

func DeleteKatakana(r *gin.Context) {
	strID := r.Param("id")
	intID, err := strconv.ParseInt(strID, 10, 64)

	if err != nil {
		utils.SendResponse(r, 400, "ID invalid", nil)
	}

	result, err := models.DeleteKatakana(intID)

	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
	}
	utils.SendResponse(r, 200, "Successfully", result)
}

func UpdateKatakana(r *gin.Context) {
	var err error

	var alphabet = models.Alphabet{}
	err = r.Bind(&alphabet)

	if err != nil || alphabet.ID == 0 {
		utils.SendResponse(r, 400, "Data invalid", nil)
	}

	models.UpdateKatakana(&alphabet)
}
