package controller

import (
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"comb.com/banking/middleware"
	"comb.com/banking/routes"
	"comb.com/banking/services"
	"github.com/labstack/echo"
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
	log.Println("Server đang chạy tại http://localhost:8080")
	if err := c.e.Start(":8080"); err != nil {
		log.Fatal("Không thể khởi động server: ", err)
	}
}

// setupRoutes khai báo các endpoint
func (c *Controller) setupRoutes() error {
	// load process
	router, err := routes.GetRouteManager()
	if err != nil {
		return err
	}
	ctrlValue := reflect.ValueOf(c.services)

	for _, r := range router.Routes {
		path := ensureSlash(r.Path)

		handlerMethod := ctrlValue.MethodByName(r.Handler)
		if !handlerMethod.IsValid() {
			log.Printf("❌ Không tìm thấy handler: %s\n", r.Handler)
			continue
		}
		// Convert reflect.Value to echo.HandlerFunc
		handlerFunc, ok := handlerMethod.Interface().(func(echo.Context) error)
		if !ok {
			log.Printf("❌ Handler %s không đúng định dạng (func(echo.Context) error)\n", r.Handler)
			continue
		}
		c.registerRoute(r.Method, path, handlerFunc, r.AuthRequired)
		log.Printf("✅ Đã thêm vào handler: %s\n", r.Handler)
	}
	return nil
}

func (c Controller) registerRoute(method, path string, handler echo.HandlerFunc, useJWT bool) {
	method = strings.ToUpper(method)

	// Chọn middleware dựa trên yêu cầu
	middlewares := []echo.MiddlewareFunc{}
	if useJWT {
		middlewares = append(middlewares, middleware.JWTMiddleware)
	}

	// Ánh xạ method HTTP với function của Echo
	switch method {
	case http.MethodGet:
		c.e.GET(path, handler, middlewares...)
	case http.MethodPost:
		c.e.POST(path, handler, middlewares...)
	case http.MethodPut:
		c.e.PUT(path, handler, middlewares...)
	case http.MethodDelete:
		c.e.DELETE(path, handler, middlewares...)
	default:
		log.Printf("❌ Method không hỗ trợ: %s\n", method)
	}
	for _, r := range c.e.Routes() {
		log.Printf("Method: %s, Path: %s, Name: %s", r.Method, r.Path, r.Name)
	}
}

func ensureSlash(path string) string {
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}
