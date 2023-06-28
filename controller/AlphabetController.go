package controller

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/cache"
	"github.com/quan12xz/basic_japanese/models"
	"github.com/quan12xz/basic_japanese/utils"
	"github.com/redis/go-redis/v9"
)

var TypeWord = map[string]int{"HIRAGANA": 1, "HIRAGANA_COMBINE": 11, "KATAKANA": 2, "KATAKANA_COMBINE": 22, "KANJI": 3}
var (
	ALPHABET_CACHE_KEY_BY_TYPE = "alphabet_cache_key_%s_%s"
)

/*
*

	Key caching:
*/
func GetAllByType(r *gin.Context) {
	strType := strings.ToUpper(r.Param("type"))
	intType, ok := TypeWord[strType]

	if !ok {
		utils.SendResponse(r, 400, "Type not exist", nil)
		return
	}

	// handle cache
	var alphabetRedisSetup = cache.RedisSetup{
		ExpireTime: time.Duration(time.Minute),
	}
	key, _ := cache.GenerateKey(r.Request.URL.Path)
	alphabetRedisSetup.Key = key

	var list *[]models.Alphabet = &[]models.Alphabet{}
	err := alphabetRedisSetup.GetData(list)
	// if data have in cache
	if err != nil && err != redis.Nil {
		//test
		// var res string
		// res = alphabetRedisSetup.GetDataTest()
		// utils.SendResponse(r, 306, "Successfully", res)
		// return
		utils.SendResponse(r, 305, "Successfully", err.Error())
		return
	} else if err == nil {
		utils.SendResponse(r, 200, "Successfully", list, true)
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

	if err != nil {
		utils.SendResponse(r, 400, "ID invalid", nil)
		return
	}

	//Generate key for process caching
	key, _ := cache.GenerateKey(r.Request.URL.Path, strID)

	// Lấy dữ liệu trong cache
	var alphabetRedisSetup = cache.RedisSetup{
		ExpireTime: time.Duration(time.Minute),
		Key:        key,
	}

	item := &models.Alphabet{}
	err = alphabetRedisSetup.GetData(item)
	if err == nil {
		utils.SendResponse(r, 200, "Successfully", item, true)
		return
	}

	// Get data from database if not cache yet
	result, err := models.GetWordByID(intID)

	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}

	// caching data if success
	alphabetRedisSetup.Value = result
	alphabetRedisSetup.SetData()
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
