package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/cache"
	"github.com/quan12xz/basic_japanese/models"
	"github.com/quan12xz/basic_japanese/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

var TypeWord = map[string]int{"HIRAGANA": 1, "HIRAGANA_COMBINE": 11, "KATAKANA": 2, "KATAKANA_COMBINE": 22, "KANJI": 3}
var alphabetRedisSetup = cache.RedisSetup{
	Context:    context.Background(),
	ExpireTime: time.Duration(time.Minute),
}

func GetAllByType(r *gin.Context) {
	strType := strings.ToUpper(r.Param("type"))
	intType, ok := TypeWord[strType]

	if !ok {
		utils.SendResponse(r, 400, "Type not exist", nil)
		return
	}

	// handle cache
	key, _ := cache.GenerateKey(r.Request.URL.Path, strType)
	alphabetRedisSetup.Key = key
	var list *[]models.Alphabet
	err := alphabetRedisSetup.GetData(list)
	// if data have in cache
	if err == nil {
		utils.SendResponse(r, 305, "Successfully", list)
		return
	}

	result, err := models.GetWordByType(intType)
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}

	//Caching data
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	alphabetRedisSetup.Value = data
	err = alphabetRedisSetup.SetData()
	if err != nil {
		utils.SendResponse(r, 200, "Successfully but not caching", err.Error())
		return
	}
	utils.SendResponse(r, 200, "Successfully", result)
}

func GetByID(r *gin.Context) {
	strID := r.Param("id")
	intID, err := strconv.ParseInt(strID, 10, 64)
	fmt.Print("successfully")
	if err != nil {
		utils.SendResponse(r, 400, "ID invalid", nil)
		return
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

func CheckCache(c *gin.Context) {
	cache.RedisSettup()
	stt := cache.RedisClient.Set(context.Background(), "test", "successfully", time.Duration(time.Minute))
	_, err := stt.Result()
	if err != nil {
		utils.SendResponse(c, 400, "Unsuccessfully", err.Error())
		return
	}
	utils.SendResponse(c, 200, "Successfully.", c.Request.URL.Path)
}
