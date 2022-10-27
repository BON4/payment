package csvupload

import (
	"context"
	"database/sql"
	"fmt"
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
	db *sql.DB
}

func NewCSVUploader[T any](db *sql.DB) *CSVUploader[T] {
	return &CSVUploader[T]{
		db: db,
	}
}

// Upload - upload from csv reader to db. Where tag - struct field tag for db column names
func (c *CSVUploader[T]) Upload(ctx context.Context, tag string, source CopyFromSource) (int64, error) {
	// file, err := os.OpenFile(filepath, os.O_RDONLY, 0777)
	// if err != nil {
	// 	return err
	// }

	names := make([]string, 0, 1)

	var t T
	for i, field := range structs.Fields(t) {
		names = append(names, field.Tag("boil"))
		fmt.Printf("%d: %s\n", i, field.Name())
	}

	tx, err := c.db.Begin()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare(pq.CopyIn("transacton_history", names...))
	if err != nil {
		panic(err)
	}

	var counter int64
	for ; source.Next(); counter++ {
		row, err := source.Values()
		fmt.Println(row, err)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		_, err = stmt.Exec(row...)
		if err != nil {
			panic(err)
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
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	return counter, nil
}
