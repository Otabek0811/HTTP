package postgres

import (
	"database/sql"
	"fmt"

	"app/config"
	_ "github.com/lib/pq"
	"app/storage"
)

type store struct{
	db *sql.DB
	category *categoryRepo
	product *productRepo
}



func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	conn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDB,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &store{
		db: db,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}



func (s *store) Category() storage.CategoryRepoI{
	if s.category == nil{
		s.category=NewCategoryRepo(s.db)
	}
	return s.category
}

func (s *store)Product() storage.ProductRepoI{
	if s.product == nil{
		s.product=NewProductRepo(s.db)
	}
	return s.product
}



