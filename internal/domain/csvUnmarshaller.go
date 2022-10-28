package domain

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/BON4/payment/internal/pkg/csvupload"
	"github.com/gocarina/gocsv"
)

type Unmarshaller struct {
	u   *gocsv.Unmarshaller
	err error
}

func NewTxHitoryUnmarshaller(f *os.File) (csvupload.CopyFromSource, error) {
	r := csv.NewReader(f)
	u, err := gocsv.NewUnmarshaller(r, TransactonHistory{})
	if err != nil {
		return nil, err
	}

	return &Unmarshaller{
		u: u,
	}, nil
}

func (p *Unmarshaller) Next() bool {
	return p.err == nil
}

// Values returns the values for the current row.
func (p *Unmarshaller) Values() ([]any, error) {
	anymodel, err := p.u.Read()
	if err != nil {
		p.err = err
		return nil, err
	}

	t, ok := anymodel.(TransactonHistory)
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
func (p *Unmarshaller) Err() error {
	return p.err
}
