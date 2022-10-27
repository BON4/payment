package pgxboiler

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PgxContextExecutor struct {
	conn *pgx.Conn
}

func NewPgxToBoilerConnection(conn *pgx.Conn) boil.ContextExecutor {
	return &PgxContextExecutor{
		conn: conn,
	}
}

func (p *PgxContextExecutor) Exec(query string, args ...interface{}) (sql.Result, error) {
	tag, err := p.conn.Exec(context.Background(), query, args...)
	p.conn.CopyFrom(
}

func (p *PgxContextExecutor) Query(query string, args ...interface{}) (*sql.Rows, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PgxContextExecutor) QueryRow(query string, args ...interface{}) *sql.Row {
	panic("not implemented") // TODO: Implement
}

func (p *PgxContextExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PgxContextExecutor) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PgxContextExecutor) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	panic("not implemented") // TODO: Implement
}
