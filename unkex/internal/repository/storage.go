package repository

import (
	
	"path/filepath"
	"runtime"
	"database/sql"
	"errors"
	
	_ "github.com/lib/pq"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4"
)


type Storage interface {
	RunMigrations(connString string) error
	Subscribe(request api.NewUserRequest) error
}

type storage struct {
	db *sql.DB
}


func NewStorage(db *sql.DB) Storage {
	return &storage {
		db: db,
	}
}


func (s *storage) RunMigrations(connString string) error {
	if connString == "" {
		return errors.New("Respository: Connection string is empty")
	}

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")
	migrationsPath := filepath.Join("file://", basePath, "/internal/repository/migrations/")

	m, err := migrate.New(migrationsPath, connString)

	if err != nil {
		return err
	}

	err = m.Up()

	return nil
}

func (s *storage) Subscribe(request api.NewUserRequest) error {
	newUserStatement := `
		INSERT INTO "user" (name, age, height, sex, activity_level, email, weight_goal) 
		VALUES ($1, $2, $3, $4, $5, $6, $7);
		`

	err := s.db.QueryRow(newUserStatement, request.Name, request.Age, request.Height, request.Sex, request.ActivityLevel, request.Email, request.WeightGoal).Err()

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}

	return nil
}