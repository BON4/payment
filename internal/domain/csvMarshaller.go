package domain

import (
	"database/sql"
	"io"

	"github.com/BON4/payment/internal/pkg/csvdownload"
	"github.com/gocarina/gocsv"
)

type Marshaller struct {
	row         *TransactonHistory
	w           gocsv.CSVWriter
	columnCount int64
	columnSet   bool
}

func NewTxHitoryMarshaller(w io.Writer) csvdownload.CopyToSource {
	return &Marshaller{
		row:         &TransactonHistory{},
		w:           gocsv.DefaultCSVWriter(w),
		columnCount: 21,
	}
}

func (m *Marshaller) Done() {
	m.w.Flush()
}

// TODO: pass here perpared []string for fewer allocs
func (m *Marshaller) Values(rows *sql.Rows) error {
	if !m.columnSet {
		names, err := rows.Columns()
		if err != nil {
			return err
		}

		if err := m.w.Write(names); err != nil {
			return err
		}
		m.columnSet = true
	}

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
		return err
	}
	m.row.MarshalCSV(s)
	return m.w.Write(s)
}
