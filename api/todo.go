package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// TodoList ...
type TodoList struct {
	GoalID           string `json:"goal_id"`
	RequiredElements string `json:"required_elements"`
	Todo             string `json:"todo"`
	SpecificGoal     string `json:"specific_goal"`
	LimitDate        string `json:"limit_date"`
}

// GetTodoList ...
func (h *Handler) GetTodoList(c *gin.Context) {
	r := models.NewTodoListRepository()
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	monthlyPlan := r.GetOne(id)

	c.JSON(http.StatusOK, gin.H{
		"monthly_plan": monthlyPlan,
	})
}

// AddTodoList ...
func (h *Handler) AddTodoList(c *gin.Context) {
	var apiTodoList TodoList
	_ = c.BindJSON(&apiTodoList)

	r := models.NewTodoListRepository()
	r.GoalID, _ = strconv.Atoi(apiTodoList.GoalID)
	r.RequiredElements = apiTodoList.RequiredElements
	r.RequiredElements = apiTodoList.RequiredElements
	r.Todo = apiTodoList.Todo
	r.SpecificGoal = apiTodoList.SpecificGoal
	r.LimitDate, _ = time.Parse("2006-01-02", apiTodoList.LimitDate)
	r.Add(&r)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   r.ID,
	})
}
