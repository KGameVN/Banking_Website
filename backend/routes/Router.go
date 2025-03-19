package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type RouteConfig struct {
	Method       string                 `json:"method"`
	Path         string                 `json:"path"`
	Handler      string                 `json:"handler"`
	Body         map[string]interface{} `json:"body,omitempty"` // optional
	AuthRequired bool                   `json:"auth_required,omitempty"`
	Description  string                 `json:"description,omitempty"`
}

type RouteManager struct {
	Routes []RouteConfig
}

var (
	once     sync.Once
	instance *RouteManager
	loadErr  error
)

// Singleton lấy ra RouteManager
func GetRouteManager() (*RouteManager, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Không thể lấy thư mục làm việc hiện tại: %v", err)
	}
	fmt.Println("Current working directory:", cwd)
	once.Do(func() {
		_ = godotenv.Load()

		routesPath := os.Getenv("ROUTE_PATH")
		if routesPath == "" {
			routesPath = "./config/routes.json" // fallback
		}

		routes, err := loadRoutesFromFile(routesPath)
		if err != nil {
			loadErr = err
			log.Println(err)
			return
		}

		instance = &RouteManager{
			Routes: routes,
		}
	})
	return instance, loadErr
}

// Đọc file và parse route
func loadRoutesFromFile(filePath string) ([]RouteConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("lỗi khi đọc file JSON: %w", err)
	}

	var routes []RouteConfig
	err = json.Unmarshal(data, &routes)
	if err != nil {
		return nil, fmt.Errorf("lỗi khi parse JSON: %w", err)
	}

	return routes, nil
}
