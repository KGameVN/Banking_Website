package controller

import (
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"

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
		} else {
			log.Printf("✅ Đã thêm vào handler: %s\n", r.Handler)
		}

		// Convert reflect.Value to echo.HandlerFunc
		handlerFunc, ok := handlerMethod.Interface().(func(echo.Context) error)
		if !ok {
			log.Printf("❌ Handler %s không đúng định dạng (func(echo.Context) error)\n", r.Handler)
			continue
		}

		// path := ensureSlash(r.Path)
		switch strings.ToUpper(r.Method) {
		case http.MethodGet:
			c.e.GET(path, handlerFunc)
		case http.MethodPost:
			c.e.POST(path, handlerFunc)
		case http.MethodPut:
			c.e.PUT(path, handlerFunc)
		case http.MethodDelete:
			c.e.DELETE(path, handlerFunc)
		default:
			log.Printf("❌ Method không hỗ trợ: %s\n", r.Method)
		}
	}
	for _, r := range c.e.Routes() {
		log.Printf("Method: %s, Path: %s, Name: %s", r.Method, r.Path, r.Name)
	}
	return nil
}

func ensureSlash(path string) string {
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}
