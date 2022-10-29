package csvdownload

import (
	"context"
	"database/sql"

	"github.com/fatih/structs"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

type CopyToSource interface {
	//Takes rows, and marshels them in to writer
	Values(*sql.Rows) error
	//Call done when there is no more rows
	Done()
}

type CSVDownloader[T any] struct {
	db *sql.DB

	tableName string

	names []string
}

func NewCSVDownloader[T any](db *sql.DB, tableName, tag string) *CSVDownloader[T] {
	names := make([]string, 0, 1)

	var t T
	for _, field := range structs.Fields(t) {
		names = append(names, field.Tag(tag))
	}

	return &CSVDownloader[T]{
		db:        db,
		tableName: tableName,
		names:     names,
	}
}

func (c *CSVDownloader[T]) Download(ctx context.Context, q *queries.Query, source CopyToSource) error {
	tx, err := c.db.Begin()

	if err != nil {
		return err

	}

	rows, err := q.QueryContext(ctx, tx)
	if err != nil {
		return err
	}

	for rows.Next() && ctx.Err() == nil {
		err := source.Values(rows)
		if err != nil {
			return err
		}
	}

	source.Done()

	err = rows.Close()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
