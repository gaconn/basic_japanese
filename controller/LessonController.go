package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quan12xz/basic_japanese/models"
	"github.com/quan12xz/basic_japanese/utils"
	"strconv"
)

func GetLessons(r *gin.Context) {
	intPage, err := strconv.Atoi(r.Query("page"))
	if err != nil {
		utils.SendResponse(r, 400, "Page invalid", nil)
		return
	}

	var l = &models.Lesson{}
	list, err := l.GetLessons(intPage)
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", err.Error())
		return
	}

	utils.SendResponse(r, 200, "Successfully", list)
}

func GetLesson(r *gin.Context) {
	intID, err := strconv.Atoi(r.Param("id"))

	if err != nil {
		utils.SendResponse(r, 400, "Page invalid", nil)
		return
	}

	var l = &models.Lesson{}
	err = l.GetLesson(intID)
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}

	utils.SendResponse(r, 200, "Successfully", l)
}

func AddLesson(r *gin.Context) {
	var l = &models.Lesson{}
	err := r.Bind(l)

	if err != nil {
		utils.SendResponse(r, 400, "Input data invalid", err.Error())
		return
	}

	err = l.AddLesson()
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}
	utils.SendResponse(r, 200, "Successfully", l)
}

func UpdateLesson(r *gin.Context) {
	var l = &models.Lesson{}
	err := r.Bind(l)

	if err != nil {
		utils.SendResponse(r, 400, "Input data invalid", nil)
		return
	}

	err = l.UpdateLessons(l.ID)
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}
	utils.SendResponse(r, 200, "Successfully", l)
}

func DeleteLesson(r *gin.Context) {
	intID, err := strconv.Atoi(r.Param("id"))

	if err != nil {
		utils.SendResponse(r, 400, "Page invalid", nil)
		return
	}

	var l = &models.Lesson{}
	l.ID = uint(intID)
	list := []*models.Lesson{l}
	err = l.DeleteLessons(list)
	if err != nil {
		utils.SendResponse(r, 400, "Unsuccessfully", nil)
		return
	}

	utils.SendResponse(r, 200, "Successfully", l)
}
