package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// GetAllDailyKpts ...
func (h *Handler) GetAllDailyKpts(c *gin.Context) {
	r := models.NewDailyKptRepository()
	dailyKpts := r.GetAll()

	c.HTML(http.StatusOK, "daily_kpts.html", gin.H{
		"dailyKpts": dailyKpts,
	})
}

// AddDailyKpt ...
func (h *Handler) AddDailyKpt(c *gin.Context) {
	loginUserID := GetLoginUserID(c)

	r := models.NewDailyKptRepository()
	r.UserID = loginUserID
	r.Keep, _ = c.GetPostForm("keep")
	r.Problem, _ = c.GetPostForm("problem")
	r.Try, _ = c.GetPostForm("try")

	r.Add(&r)
	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// GetDailyKpt ...
func (h *Handler) GetDailyKpt(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	dailyKpt := r.GetOne(id)

	c.HTML(http.StatusOK, "daily_kpt_edit.html", gin.H{
		"dailyKpt": dailyKpt,
	})
}

// EditDailyKpt ...
func (h *Handler) EditDailyKpt(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	dailyKpt := r.GetOne(id)

	dailyKpt.Keep, _ = c.GetPostForm("keep")
	dailyKpt.Problem, _ = c.GetPostForm("problem")
	dailyKpt.Try, _ = c.GetPostForm("try")

	r.Edit(dailyKpt)
	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// IncreaseGood ...
func (h *Handler) IncreaseGood(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	dailyKpt := r.GetOne(id)
	dailyKpt.Good++
	r.Edit(dailyKpt)

	loginUserID := GetLoginUserID(c)

	rah := models.NewKptReactionHistoryRepository()
	rah.AddReaction(int(dailyKpt.ID), loginUserID, models.ReactionGood)

	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// IncreaseFight ...
func (h *Handler) IncreaseFight(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	dailyKpt := r.GetOne(id)
	dailyKpt.Fight++
	r.Edit(dailyKpt)

	loginUserID := GetLoginUserID(c)

	rah := models.NewKptReactionHistoryRepository()
	rah.AddReaction(int(dailyKpt.ID), loginUserID, models.ReactionFight)

	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// DeleteDailyKpt ...
func (h *Handler) DeleteDailyKpt(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}
