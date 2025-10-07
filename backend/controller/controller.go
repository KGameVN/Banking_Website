package controller

import (
	"log"
	"sync"

	appmid "comb.com/banking/core/middleware"
	"comb.com/banking/internal/services"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	e        *echo.Echo
	services *services.Service
}

var (
	instance *Controller
	once     sync.Once
)

// GetController khởi tạo singleton cho Controller
func GetController() *Controller {
	once.Do(func() {
		e := echo.New()

		// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// 	AllowOrigins: []string{"http://localhost:3000"},
		// 	AllowMethods: []string{
		// 		http.MethodGet,
		// 		http.MethodPost,
		// 		http.MethodPut,
		// 		http.MethodDelete,
		// 		http.MethodOptions,
		// 	},
		// }))

		svc := services.GetService()
		if svc == nil {
			log.Fatal("Không thể khởi tạo service")
			return
		}

		instance = &Controller{
			e:        e,
			services: svc,
		}

		instance.setupRoutes()
	})

	return instance
}

// Start khởi động server
func (c *Controller) Start() {
	// certFile := "certs/cert.pem"
	// keyFile := "certs/key.pem"

	// log.Println("Server đang chạy tại https://localhost:8443")
	// if err := c.e.StartTLS(":8443", certFile, keyFile); err != nil {
	// 	log.Fatal("Không thể khởi động server HTTPS: ", err)
	// }
	c.e.Use(mid.CORS())
	log.Println("Server đang chạy tại http://localhost:8080")
	if err := c.e.Start(":8080"); err != nil {
		log.Fatal("Không thể khởi động server: ", err)
	}
}

// setupRoutes khai báo các endpoint
func (c *Controller) setupRoutes() error {
	// main group - error handler
	apiGroup := c.e.Group("/api", appmid.ErrorHandlerMiddleware)

	// register
	apiGroup.POST("/register", c.services.Register)

	// login - jwt
	userGroup := apiGroup.Group("/user", appmid.JWTMiddleware)
	userGroup.POST("/login", c.services.Login)

	// account - jwt
	accountGroup := apiGroup.Group("/account", appmid.JWTMiddleware)
	accountGroup.GET("/:id", c.services.GetAccountInfo)
	accountGroup.POST("/transaction/:account", c.services.Transaction)
	accountGroup.GET("/transaction/:account/history", c.services.GetTransHistory)
	accountGroup.POST("/transfer", c.services.Transfer)
	return nil
}
