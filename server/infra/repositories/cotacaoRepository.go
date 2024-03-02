package repositories

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

func (s *ServerRepository) CreateTable() error {
	sqlStmt := `
	CREATE table cotacoes(coin text, bid text);
	DELETE from cotacoes;
	`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}

func (s *ServerRepository) Insert(ctx context.Context, coin, bid string) error {
	tx, err := s.db.Begin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ao iniciar transaction %s\n", err)
		return err
	}

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

	err = tx.Commit()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executar commit %s\n", err)
		return err
	}

	return nil
}
