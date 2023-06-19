package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/quan12xz/basic_japanese/models"
	"github.com/quan12xz/basic_japanese/utils"
)

func GetKatakana(r *gin.Context) {
	utils.SendResponse(r, 200, "Successfully", nil)
}

func AddKatakana(r *gin.Context) {
	var alphabet models.Alphabet

	json.Unmarshal([]byte(r.Request.Body), &alphabet)
}
