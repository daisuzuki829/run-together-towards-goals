package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID             string `json:"id"`
	Nickname       string `json:"nickname"`
	Password       string `json:"password"`
	Age            string `json:"age"`
	AgeDisplayFlag string `json:"age_display_flag"`
	Address        string `json:"address"`
	BirthPlace     string `json:"birth_place"`
	Hobby          string `json:"hobby"`
	Occupation     string `json:"occupation"`
	StrongPoint    string `json:"strong_point"`
	Skill          string `json:"skill"`
	Role           string `json:"role"`
}

// GetUser ...
func (h *Handler) GetUser(c *gin.Context) {
	r := models.NewUserRepository()
	userID, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	user := r.GetOne(userID)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// AddUser ...
func (h *Handler) AddUser(c *gin.Context) {
	var apiUser User
	_ = c.BindJSON(&apiUser)

	var user models.User
	user.Nickname = apiUser.Nickname
	user.Password = apiUser.Password
	user.Age = apiUser.Age
	user.AgeDisplayFlag, _ = strconv.Atoi(apiUser.AgeDisplayFlag)
	user.Address = apiUser.Address
	user.BirthPlace = apiUser.BirthPlace
	user.Hobby = apiUser.Hobby
	user.Occupation = apiUser.Occupation
	user.StrongPoint = apiUser.StrongPoint
	user.Skill = apiUser.Skill
	user.Role, _ = strconv.Atoi(apiUser.Role)

	r := models.NewUserRepository()
	err := r.Add(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"id":   user.ID,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   user.ID,
	})
}

// EditUser ...
func (h *Handler) EditUser(c *gin.Context) {
	var beforeUser User
	c.BindJSON(&beforeUser)

	id, _ := strconv.Atoi(beforeUser.ID)
	r := models.NewUserRepository()
	user := r.GetOne(id)

	user.Nickname = beforeUser.Nickname
	if beforeUser.Password == "" {
		password, _ := bcrypt.GenerateFromPassword([]byte(beforeUser.Password), bcrypt.DefaultCost)
		user.Password = string(password)
	}
	user.Age = beforeUser.Age
	user.AgeDisplayFlag, _ = strconv.Atoi(beforeUser.AgeDisplayFlag)
	user.Address = beforeUser.Address
	user.BirthPlace = beforeUser.BirthPlace
	user.Hobby = beforeUser.Hobby
	user.Occupation = beforeUser.Occupation
	user.StrongPoint = beforeUser.StrongPoint
	user.Skill = beforeUser.Skill

	err := r.Edit(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"id":   user.ID,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   user.ID,
	})
}
