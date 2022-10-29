package csvupload

import (
	"context"
	"database/sql"
	"io"

	"github.com/fatih/structs"
	"github.com/lib/pq"
)

type CopyFromSource interface {
	// Next returns true if there is another row and makes the next row data
	// available to Values(). When there are no more rows available or an error
	// has occurred it returns false.
	Next() bool

	// Values returns the values for the current row.
	Values() ([]any, error)

	Err() error
}

type CSVUploader[T any] struct {
	db        *sql.DB
	tableName string
	tag       string
	names     []string
}

func NewCSVUploader[T any](db *sql.DB, tableName, tag string) *CSVUploader[T] {
	names := make([]string, 0, 1)

	var t T
	for _, field := range structs.Fields(t) {
		names = append(names, field.Tag("boil"))
	}

	return &CSVUploader[T]{
		db:        db,
		tableName: tableName,
		names:     names,
		tag:       tag,
	}
}

// Upload - upload from csv reader to db. Where tag - struct field tag for db column names
func (c *CSVUploader[T]) Upload(ctx context.Context, source CopyFromSource) (int64, error) {
	// file, err := os.OpenFile(filepath, os.O_RDONLY, 0777)
	// if err != nil {
	// 	return err
	// }

	tx, err := c.db.Begin()

	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(pq.CopyIn(c.tableName, c.names...))
	if err != nil {
		return 0, err
	}

	var counter int64
	for ; source.Next() && ctx.Err() == nil; counter++ {
		row, err := source.Values()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		_, err = stmt.Exec(row...)
		if err != nil {
			return 0, err
		}
	}

	if err := source.Err(); err != nil && err != io.EOF {
		return 0, err
	}

	// res, err := stmt.Exec()
	// if err != nil {
	// 	panic(err)
	// }

	// n, err := res.RowsAffected()
	// if err != nil {
	// 	return 0, err
	// }

	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return counter, nil
}
