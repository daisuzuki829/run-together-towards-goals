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
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Role     string `json:"role"`
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
	user.Role,_ = strconv.Atoi(apiUser.Role)

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
