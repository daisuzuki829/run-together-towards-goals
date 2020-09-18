package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"github.com/jinzhu/gorm"
)

// Goal ...
type Goal struct {
	gorm.Model
	UserID      string `json:"user_id"`
	GenreID     string `json:"genre_id"`
	GoalName    string `json:"goal_name"`
	DisplayFlag string `json:"display_flag"`
}

// SetMyGoal ...
func (h *Handler) SetMyGoal(c *gin.Context) {
	var apiMyGoal Goal
	_ = c.BindJSON(&apiMyGoal)

	r := models.NewGoalRepository()
	r.UserID, _ = strconv.Atoi(apiMyGoal.UserID)
	r.GenreID, _ = strconv.Atoi(apiMyGoal.GenreID)
	r.GoalName = apiMyGoal.GoalName
	r.DisplayFlag, _ = strconv.Atoi(apiMyGoal.DisplayFlag)
	r.Add(&r)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   r.ID,
	})
}

// EditMyGoal ...
func (h *Handler) EditMyGoal(c *gin.Context) {
	r := models.NewGoalRepository()

	id := c.DefaultQuery("id", "0")
	myGoalID, _ := strconv.Atoi(id)
	newMyGoal := r.GetOne(myGoalID)

	if newMyGoal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "Not Found",
		})
		return
	}

	var goal Goal
	_ = c.BindJSON(&goal)

	newMyGoal.GenreID, _ = strconv.Atoi(goal.GenreID)
	newMyGoal.GoalName = goal.GoalName
	newMyGoal.DisplayFlag, _ = strconv.Atoi(goal.DisplayFlag)

	r.Edit(newMyGoal)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   newMyGoal.ID,
	})
}
