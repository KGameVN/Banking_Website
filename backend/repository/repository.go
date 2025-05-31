package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"

	"comb.com/banking/ent"
)

type Repository struct {
	DbClient *ent.Client
}

type RepositoryEnv struct {
	DBUser     string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

var env = &RepositoryEnv{}

var (
	instance *Repository
	once     sync.Once
)

// Hàm khởi tạo singleton
func GetRepository() *Repository {
	once.Do(func() {
		repo, err := initRepository()
		if err != nil {
			log.Println("Can't connect DB")
		}
		instance = repo
	})
	return instance
}

func initRepository() (*Repository, error) {
	if err := loadRepositoryEnv(); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("load env error: %w")
	}
	dns := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	log.Println(dns)
	client, err := loadClient(dns)
	if err != nil {
		return nil, fmt.Errorf("ent client error : %w", err)
	}
	return &Repository{DbClient: client}, nil
}

// LoadRepositoryEnv đọc file .env và gán vào biến `env`
func loadRepositoryEnv() error {
	_ = godotenv.Load("../.env") // ignore error in production

	envget := func(key string, required bool) string {
		val, exists := os.LookupEnv(key)
		if !exists && required {
			log.Fatalf("Missing required env: %s", key)
		}
		return val
	}

	env.DBUser = envget("DB_USER", true)
	env.DBPassword = envget("DB_PASSWORD", true)
	env.DBPort = envget("DB_PORT", true)
	env.DBName = envget("DB_NAME", true)
	env.DBPort = envget("DB_PORT", true)

	if os.Getenv("ENV") == "dev" {
		env.DBHost = os.Getenv("DB_HOST_DEV")
	} else {
		env.DBHost = os.Getenv("DB_HOST_PROD")
	}
	return nil
}

// GetRepositoryEnv trả về con trỏ đến `env`
func (r Repository) GetRepositoryEnv() *RepositoryEnv {
	return env
}

func loadClient(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
