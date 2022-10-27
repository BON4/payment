package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/BON4/payment/internal/csvupload"
	"github.com/BON4/payment/internal/domain"
	"github.com/gocarina/gocsv"

	_ "github.com/lib/pq"
)

var filepath = "example.csv"

type Parser struct {
	u   *gocsv.Unmarshaller
	err error
}

func NewTxHitoryParser(f *os.File) (csvupload.CopyFromSource, error) {

	r := csv.NewReader(f)
	u, err := gocsv.NewUnmarshaller(r, domain.TransactonHistory{})
	if err != nil {
		return nil, err
	}

	return &Parser{
		u: u,
	}, nil
}

func (p *Parser) Next() bool {
	return p.err == nil
}

// Values returns the values for the current row.
func (p *Parser) Values() ([]any, error) {
	anymodel, err := p.u.Read()
	if err != nil {
		p.err = err
		return nil, err
	}

	t, ok := anymodel.(domain.TransactonHistory)
	if ok {

		arr := []interface{}{
			t.TransactionId,
			t.RequestId,
			t.TerminalId,
			t.PartnerObjectId,
			t.AmountTotal,
			t.AmountOriginal,
			t.CommissionPS,
			t.CommissionClient,
			t.CommissionProvider,
			t.DateInput,
			t.DatePost,
			t.Status,
			t.PaymentType,
			t.PaymentNumber,
			t.ServiceId,
			t.Service,
			t.PayeeId,
			t.PayeeName,
			t.PayeeBankMfo,
			t.PayeeBankAccount,
			t.PaymentNarrative,
		}

		return arr, nil
	}

	p.err = errors.New("invalid type")

	return nil, p.err
}

// Err returns any error that has been encountered by the CopyFromSource. If
// this is not nil *Conn.CopyFrom will abort the copy.
func (p *Parser) Err() error {
	return p.err
}

func main() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/payment?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	uploader := csvupload.NewCSVUploader[domain.TransactonHistory](conn)

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	par, err := NewTxHitoryParser(file)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	n, err := uploader.Upload(context.Background(), par)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed %d objects from csv", n)
}
