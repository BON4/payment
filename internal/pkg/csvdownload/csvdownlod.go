package csvdownload

import (
	"context"
	"database/sql"
	"io"

	"github.com/fatih/structs"
	"github.com/gocarina/gocsv"
)

type CopyToSource interface {
	//Takes rows, and returns parsed strings
	Values(*sql.Rows) ([]string, error)
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

func (c *CSVDownloader[T]) Download(ctx context.Context, out io.Writer, source CopyToSource) error {
	w := gocsv.DefaultCSVWriter(out)

	tx, err := c.db.Begin()

	if err != nil {
		return err
	}

	rows, err := tx.QueryContext(ctx, "select * from "+c.tableName)
	if err != nil {
		return err
	}

	names, err := rows.Columns()
	if err != nil {
		return err
	}

	err = w.Write(names)
	if err != nil {
		return err
	}

	for rows.Next() && ctx.Err() == nil {
		row, err := source.Values(rows)
		if err != nil {
			return err
		}

		err = w.Write(row)
		if err != nil {
			return err
		}
	}

	w.Flush()

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
