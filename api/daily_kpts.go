package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// DailyKpt ...
type DailyKpt struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Keep    string `json:"keep"`
	Problem string `json:"problem"`
	Try     string `json:"try"`
}

// PostDailyKpt ...
func (h *Handler) PostDailyKpt(c *gin.Context) {
	var apiDailyKpt DailyKpt
	_ = c.BindJSON(&apiDailyKpt)

	r := models.NewDailyKptRepository()
	r.UserID, _ = strconv.Atoi(apiDailyKpt.UserID)
	r.Keep = apiDailyKpt.Problem
	r.Problem = apiDailyKpt.Problem
	r.Try = apiDailyKpt.Try

	if err := r.Add(&r); err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   r.ID,
	})
}

// EditDailyKpt ...
func (h *Handler) EditDailyKpt(c *gin.Context) {
	var apiDailyKpt DailyKpt
	c.BindJSON(&apiDailyKpt)

	id, _ := strconv.Atoi(apiDailyKpt.ID)
	r := models.NewDailyKptRepository()
	dailyKpt := r.GetOne(id)

	dailyKpt.UserID, _ = strconv.Atoi(apiDailyKpt.UserID)
	dailyKpt.Keep = apiDailyKpt.Problem
	dailyKpt.Problem = apiDailyKpt.Problem
	dailyKpt.Try = apiDailyKpt.Try

	r.Edit(dailyKpt)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Update",
		"id":   dailyKpt.ID,
	})
}

// GetDailyKpts ...
func (h *Handler) GetDailyKpts(c *gin.Context) {
	r := models.NewDailyKptRepository()

	c.JSON(http.StatusOK, gin.H{
		"daily_kpts": r.GetAll(),
	})
}

// IncreaseGood ...
func (h *Handler) IncreaseGood(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	r := models.NewDailyKptRepository()
	dailyKpt := r.GetOne(id)

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}

	dailyKpt.Good++

	r.Edit(dailyKpt)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Update",
		"id":   dailyKpt.ID,
	})
}

// IncreaseFight ...
func (h *Handler) IncreaseFight(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	r := models.NewDailyKptRepository()
	dailyKpt := r.GetOne(id)

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}

	dailyKpt.Fight++

	r.Edit(dailyKpt)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Update",
		"id":   dailyKpt.ID,
	})
}

// GetHello ...
func GetHello(userName string) string {
	return "Hello, " + userName + "!!"
}
