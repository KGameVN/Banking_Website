package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"comb.com/banking/ent"
	"comb.com/banking/ent/migrate"
	"comb.com/banking/ent/seed"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Repository struct {
	DbClient *ent.Client
}

type Env struct {
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBName     string `env:"DB_NAME" envDefault:"dbname"`
}

var env = &Env{}

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
	log.Printf(env.DBName)
	return nil
}

// GetRepositoryEnv trả về con trỏ đến `env`
func (r Repository) GetRepositoryEnv() *Env {
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

func (r Repository) GenerateSchema() error {
	ctx := context.Background()
	// Auto migration
	if err := r.DbClient.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Seed dummy data
	seed.SeedData(ctx, r.DbClient)
	log.Println("✅ Dummy data inserted successfully!")
	return nil
}
