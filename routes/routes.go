package routes

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/api"
	"github.com/hariNEzuMI928/run-together-towards-goals/controllers"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"github.com/jinzhu/gorm"
)

// Handler ...
func Handler(dbConn *gorm.DB) {

	handler := controllers.Handler{
		Db: dbConn,
	}
	apiHandler := api.Handler{
		Db: dbConn,
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "DELETE", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")

	// セッションの設定
	r.Use(sessions.Sessions("my_session", cookie.NewStore([]byte("secret"))))

	// login
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.GET("/registration", func(c *gin.Context) {
		rg := models.NewGenreRepository()
		genre := rg.GetAll()
		c.HTML(http.StatusOK, "registration.html", gin.H{
			"genre": genre,
		})
	})
	r.POST("/registration", controllers.NewRegistration)
	r.GET("/forgot_password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forgot_password.html", gin.H{
			"title": "title",
		})
	})
	r.POST("/forgot_password", controllers.ForgotPassword)

	r.Group("/")
	r.Use(controllers.SessionCheck)
	{
		// user info
		r.GET("/_users", handler.GetAllUsers)
		rUser := r.Group("/user")
		{
			rUser.GET("my_page", handler.GetMyPage)
			rUser.POST("add", handler.AddUser)
			rUser.GET("view/:id", handler.GetUser)
			rUser.POST("edit_ok/:id", handler.EditUser)
			rUser.POST("edit_my_page", handler.EditMyPage)
			rUser.POST("delete/:id", handler.DeleteUser)
			rUser.POST("user_edit_goal/:id", handler.EditGoal)
			rUser.POST("user_delete_goal/:id", handler.DeleteGoal)
			rUser.POST("user_edit_todo/:id", handler.EditTodo)
			rUser.POST("user_delete_todo/:id", handler.DeleteTodo)
			rUser.POST("add_goal", handler.AddGoal)
			rUser.POST("add_todo", handler.AddTodo)
		}

		// monthly_plans
		r.GET("/_monthly_plans", handler.GetAllMonthlyPlans)
		rMonthlyPlan := r.Group("/monthly_plan")
		{
			rMonthlyPlan.POST("add", handler.AddMonthlyPlans)
			rMonthlyPlan.GET("edit/:id", handler.GetMonthlyPlan)
			rMonthlyPlan.POST("edit_ok/:id", handler.EditMonthlyPlan)
			rMonthlyPlan.POST("delete/:id", handler.DeleteMonthlyPlan)
		}

		// daily kpt info
		r.GET("/_daily_kpts", handler.GetAllDailyKpts)
		rDailyKpt := r.Group("/daily_kpt")
		{
			rDailyKpt.POST("add", handler.AddDailyKpt)
			rDailyKpt.GET("edit/:id", handler.GetDailyKpt)
			rDailyKpt.POST("edit_ok/:id", handler.EditDailyKpt)
			rDailyKpt.POST("good/:id", handler.IncreaseGood)
			rDailyKpt.POST("fight/:id", handler.IncreaseFight)
			rDailyKpt.POST("delete/:id", handler.DeleteDailyKpt)
		}

		// genre info
		r.GET("/_genres", handler.GetAllGenres)
		rGenre := r.Group("/genre")
		{
			rGenre.POST("add", handler.AddGenre)
			rGenre.GET("edit/:id", handler.GetGenre)
			rGenre.POST("edit_ok/:id", handler.EditGenre)
			rGenre.GET("delete/:id", handler.DeleteGenre)
		}

		r.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "welcome.html", gin.H{
				"user":  controllers.GetLoginUser(c),
			})
		})

		r.NoRoute(func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/index")
		})
	}

	rAPI := r.Group("/api")
	{
		rAPIUser := rAPI.Group("/user")
		{
			rAPIUser.GET("", apiHandler.GetUser)
			rAPIUser.POST("add", apiHandler.AddUser)
			rAPIUser.PUT("edit", apiHandler.EditUser)
		}
		rMonthlyPlan := rAPI.Group("/monthly_plan")
		{
			rMonthlyPlan.GET("", apiHandler.GetMonthlyPlan)
			rMonthlyPlan.POST("add", apiHandler.AddMonthlyPlan)
			rMonthlyPlan.PUT("add", apiHandler.EditMonthlyPlan)
		}
		rAPIDailyKpt := rAPI.Group("/daily_kpt")
		{
			rAPIDailyKpt.GET("", apiHandler.GetDailyKpts)
			rAPIDailyKpt.POST("add", apiHandler.PostDailyKpt)
			rAPIDailyKpt.PUT("good", apiHandler.IncreaseGood)
			rAPIDailyKpt.PUT("fight", apiHandler.IncreaseFight)
		}
		rAPIMyGoal := rAPI.Group("/my_goals")
		{
			rAPIMyGoal.POST("add", apiHandler.SetMyGoal)
			rAPIMyGoal.PUT("edit", apiHandler.EditMyGoal)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		// TODO portが既に使われているときの対応
		log.Println(err.Error())
		os.Exit(1)
	}
}
