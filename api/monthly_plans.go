package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// MonthlyPlan ...
type MonthlyPlan struct {
	ID                 string `json:"id"`
	UserID             string `json:"user_id"`
	GoalID             string `json:"goal_id"`
	Month              string `json:"month"`
	KeepInLastMonth    string `json:"keep_in_last_month"`
	ProblemInLastMonth string `json:"problem_in_last_month"`
	GoalAfterHalfYear  string `json:"goal_after_half_year"`
	GoalInThisMonth    string `json:"goal_in_this_month"`
	CurrentState       string `json:"current_state"`
	DailyTodo          string `json:"daily_todo"`
}

// GetMonthlyPlan ...
func (h *Handler) GetMonthlyPlan(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	monthlyPlan := r.GetOne(id)

	c.JSON(http.StatusOK, gin.H{
		"monthly_plan": monthlyPlan,
	})
}

// AddMonthlyPlan ...
func (h *Handler) AddMonthlyPlan(c *gin.Context) {
	var monthlyPlan MonthlyPlan
	_ = c.BindJSON(&monthlyPlan)

	r := models.NewMonthlyPlanRepository()
	r.UserID, _ = strconv.Atoi(monthlyPlan.UserID)
	r.GoalID, _ = strconv.Atoi(monthlyPlan.GoalID)
	r.Month, _ = time.Parse("2006-01", monthlyPlan.Month)
	r.KeepInLastMonth = monthlyPlan.KeepInLastMonth
	r.ProblemInLastMonth = monthlyPlan.ProblemInLastMonth
	r.GoalAfterHalfYear = monthlyPlan.GoalAfterHalfYear
	r.GoalInThisMonth = monthlyPlan.GoalInThisMonth
	r.CurrentState = monthlyPlan.CurrentState
	r.DailyTodo = monthlyPlan.DailyTodo

	r.Add(&r)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   r.ID,
	})
}

// EditMonthlyPlan ...
func (h *Handler) EditMonthlyPlan(c *gin.Context) {
	var apiMonthlyPlan MonthlyPlan
	c.BindJSON(&apiMonthlyPlan)

	id, _ := strconv.Atoi(apiMonthlyPlan.ID)
	r := models.NewMonthlyPlanRepository()
	monthlyPlan := r.GetOne(id)

	monthlyPlan.UserID, _ = strconv.Atoi(apiMonthlyPlan.UserID)
	monthlyPlan.GoalID, _ = strconv.Atoi(apiMonthlyPlan.GoalID)
	monthlyPlan.Month, _ = time.Parse("2006-01", apiMonthlyPlan.Month)
	monthlyPlan.KeepInLastMonth = apiMonthlyPlan.KeepInLastMonth
	monthlyPlan.ProblemInLastMonth = apiMonthlyPlan.ProblemInLastMonth
	monthlyPlan.GoalAfterHalfYear = apiMonthlyPlan.GoalAfterHalfYear
	monthlyPlan.GoalInThisMonth = apiMonthlyPlan.GoalInThisMonth
	monthlyPlan.KeepInLastMonth = apiMonthlyPlan.KeepInLastMonth
	monthlyPlan.CurrentState = apiMonthlyPlan.CurrentState
	monthlyPlan.DailyTodo = apiMonthlyPlan.DailyTodo

	r.Edit(monthlyPlan)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Update",
		"id":   monthlyPlan.ID,
	})
}

// DeletetMonthlyPlan ...
func (h *Handler) DeletetMonthlyPlan(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	r := models.NewMonthlyPlanRepository()

	r.Delete(id)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Update",
		"id":   id,
	})
}
