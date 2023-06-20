package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/models"
	"github.com/quan12xz/basic_japanese/utils"
)

func GetKatakana(r *gin.Context) {
	utils.SendResponse(r, 200, "Successfully", nil)
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
