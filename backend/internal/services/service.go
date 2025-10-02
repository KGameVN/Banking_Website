package services

import (
	"log"
	"sync"

	"comb.com/banking/internal/repository"
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
		err := repo.GenerateSchema()
		if err != nil {
			log.Println("Seed generete dummy error")
		}
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
