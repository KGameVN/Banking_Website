package controller

import (
	"log"
	"sync"

	a "comb.com/banking/middleware"
	"comb.com/banking/services"
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

	c.e.POST("/login", c.services.Login)

	userGroup := c.e.Group("/account", a.JWTMiddleware)
	userGroup.GET("/trans_history", c.services.GetTransHistory)
	userGroup.POST("/transfer", c.services.Transfer)
	return nil
}
