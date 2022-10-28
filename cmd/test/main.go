package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/BON4/payment/internal/domain"
	"github.com/BON4/payment/internal/pkg/csvdownload"
	"github.com/BON4/payment/internal/pkg/csvupload"
	"github.com/gocarina/gocsv"

	_ "github.com/lib/pq"
)

var filepath = "example.csv"

type Marshaller struct {
	row         *domain.TransactonHistory
	columnCount int64
}

func NewMarshaler() csvdownload.CopyToSource {
	return &Marshaller{
		row:         &domain.TransactonHistory{},
		columnCount: 21,
	}
}

// TODO: pass here perpared []string for fewer allocs
func (m *Marshaller) Values(rows *sql.Rows) ([]string, error) {
	s := make([]string, m.columnCount)
	if err := rows.Scan(
		&m.row.TransactionId,
		&m.row.RequestId,
		&m.row.TerminalId,
		&m.row.PartnerObjectId,
		&m.row.AmountTotal,
		&m.row.AmountOriginal,
		&m.row.CommissionPS,
		&m.row.CommissionClient,
		&m.row.CommissionProvider,
		&m.row.DateInput,
		&m.row.DatePost,
		&m.row.Status,
		&m.row.PaymentType,
		&m.row.PaymentNumber,
		&m.row.ServiceId,
		&m.row.Service,
		&m.row.PayeeId,
		&m.row.PayeeName,
		&m.row.PayeeBankMfo,
		&m.row.PayeeBankAccount,
		&m.row.PaymentNarrative,
	); err != nil {
		return s, err
	}
	m.row.MarshalCSV(s)
	return s, nil
}

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
	conn, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/payment_test?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	uploader := csvupload.NewCSVUploader[domain.TransactonHistory](conn, "transacton_history")

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	par, err := NewTxHitoryParser(file)
	if err != nil {
		panic(err)
	}

	n, err := uploader.Upload(context.Background(), "boil", par)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	file.Close()

	file, err = os.OpenFile("new_example.csv", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	marshaler := NewMarshaler()
	downloader := csvdownload.NewCSVDownloader[domain.TransactonHistory](conn, "transacton_history", "tag")

	err = downloader.Download(context.Background(), file, marshaler)
	if err != nil {
		panic(err)
	}

	file.Close()
}
