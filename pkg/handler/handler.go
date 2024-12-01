package handler

import (
	"orden/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "orden/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	/*
		filesPath, err := filepath.Abs(os.Getenv("FILES_PATH"))
		//"/root/directory/files"
		if err != nil {
			panic(err)
		}
		logrus.Println("filesPath", filesPath)
		router.Static("/files", filesPath)
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signup)
		auth.POST("/login", h.login)

		auth.GET("/steam", h.signupSteam)
		auth.GET("/steam/callback", h.callbackSteam)

		auth.PATCH("/resetpasswd", h.resetpasswd)

		auth.POST("/adminlogin", h.loginForAdmin)
	}

	api := router.Group("/api", h.userIdentity)
	{

		profile := api.Group("/profile")
		{
			profile.PATCH("/changepasswd", h.changePasswd)
			profile.PATCH("/changeusername", h.updateUsername)
			profile.GET("/", h.getProfile)
			profile.DELETE("/", h.deleteUser)

		}

		admin := api.Group("/admin", h.checkAdmin)
		{
			admin.GET("/users", h.getAllUsers)
			admin.PATCH("/adminrights/:id", h.giveAdminRight)
			admin.DELETE("/adminrights/:id", h.removeAdminRight)
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router

}
