package api

import (
	"net/http"
	"strconv"

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
