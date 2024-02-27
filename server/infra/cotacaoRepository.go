package infra

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

type ServerRepository struct {
	db *sql.DB
}

func NewServerRepository(db *sql.DB) *ServerRepository {
	return &ServerRepository{
		db: db,
	}
}

func (s *ServerRepository) Insert(ctx context.Context, coin, bid string) error {
	stmt, err := s.db.PrepareContext(ctx, "INSERT INTO cotacoes (coin, bid) values(?, ?)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao preparar query de insert %s\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, coin, bid)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao executar insert %s\n", err)
		return err
	}

	return nil
}
