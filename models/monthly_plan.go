package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// MonthlyPlan ...
type MonthlyPlan struct {
	gorm.Model
	UserID             int       `gorm:"not null" validate:"required,numeric"`
	GoalID             int       `gorm:"not null" validate:"required,numeric"`
	Month              time.Time `gorm:"not null" sql:"not null;type:date"`
	KeepInLastMonth    string
	ProblemInLastMonth string
	GoalAfterHalfYear  string
	GoalInThisMonth    string
	CurrentState       string
	DailyTodo          string
}

// MonthlyPlans ...
type MonthlyPlans struct {
	gorm.Model
	UserID             int
	Nickname           string
	GoalName           string
	GoalID             int
	Month              time.Time
	KeepInLastMonth    string
	ProblemInLastMonth string
	GoalAfterHalfYear  string
	GoalInThisMonth    string
	CurrentState       string
	DailyTodo          string
}

// NewMonthlyPlanRepository ...
func NewMonthlyPlanRepository() MonthlyPlan {
	return MonthlyPlan{}
}

// Add ...
func (o *MonthlyPlan) Add(monthlyPlan *MonthlyPlan) {
	db := Open()
	db.Create(monthlyPlan)
	defer db.Close()
}

// Edit ...
func (o *MonthlyPlan) Edit(monthlyPlan MonthlyPlan) {
	db := Open()
	db.Save(monthlyPlan)
	db.Close()
}

// GetAll ...
func (o *MonthlyPlan) GetAll(id int) []MonthlyPlans {
	db := Open()
	var monthlyPlans []MonthlyPlans
	db.Table("monthly_plans").
		Select("monthly_plans.*, users.nickname, goals.goal_name").
		Joins("inner JOIN users ON monthly_plans.user_id = users.id").
		Joins("inner JOIN goals ON monthly_plans.goal_id = goals.id").
		Where("monthly_plans.deleted_at is null").
		Where("users.deleted_at is null").
		Where("goals.deleted_at is null").
		// ログインユーザーの非表示設定目標は表示して、それ以外のユーザーの非表示目標は隠す
		Where("(users.id = ?) or (users.id <> ? and goals.display_flag <> 0)", id, id).
		Order("users.id, monthly_plans.id").
		Find(&monthlyPlans)
	db.Close()
	return monthlyPlans
}

// GetOne ...
func (o *MonthlyPlan) GetOne(id int) MonthlyPlan {
	db := Open()
	var monthlyPlan MonthlyPlan
	db.Find(&monthlyPlan, id)
	db.Close()
	return monthlyPlan
}

// Delete ...
func (o *MonthlyPlan) Delete(id int) {
	db := Open()
	var monthlyPlan MonthlyPlan
	db.First(&monthlyPlan, id)
	db.Delete(&monthlyPlan)
	db.Close()
}

// Count ...
func (o *MonthlyPlan) Count() int {
	db := Open()
	var count = 0
	db.Table("monthly_plans").Count(&count)
	db.Close()
	return count
}
