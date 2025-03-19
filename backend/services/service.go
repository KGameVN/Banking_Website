package services

import (
	"log"
	"sync"

	"comb.com/banking/repository"
)

type Service struct {
	Repository *repository.Repository
}

var (
	instance *Service
	once     sync.Once
)

// GetService khởi tạo singleton cho Service
func GetService() *Service {
	once.Do(func() {
		repo := repository.GetRepository()
		if repo == nil {
			log.Println("Cannot initialize service: repository is nil")
			return
		}
		instance = &Service{
			Repository: repo,
		}
	})
	return instance
}
