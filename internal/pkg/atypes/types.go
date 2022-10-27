package atypes

import (
	"database/sql/driver"
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
)

type DateTime struct {
	time.Time
}

// Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02 15:04:05"), nil
}

// You could also use the standard Stringer interface
func (date *DateTime) String() string {
	return date.String() // Redundant, just for example
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}

func (date DateTime) Value() (driver.Value, error) {
	return pq.FormatTimestamp(date.Time), nil
}

func (date *DateTime) Scan(src interface{}) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", string(src.([]byte)))
	return err
}

type Decimal struct {
	decimal.Decimal
}

func (d *Decimal) MarshalCSV() (string, error) {
	return d.String(), nil
}

func (d *Decimal) UnmarshalCSV(csv string) (err error) {
	d.Decimal, err = decimal.NewFromString(csv)
	return err
}
