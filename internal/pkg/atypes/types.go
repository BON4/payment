package atypes

import (
	"database/sql/driver"
	"time"

	"github.com/shopspring/decimal"
)

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

type DateTime struct {
	time.Time
}

// All this to rune conversion is just to delete trailing zero in hour
// Golang dont have specified format tor this
func (date *DateTime) MarshalCSV() (string, error) {
	//s := []rune(date.Format("2006-01-02 15:04:05"))
	//if date.Hour() < 10 {
	//	s = delChar(s, len(s)-8)
	//}
	return date.Format("2006-01-02 15:04:05"), nil
}

func (date *DateTime) MarshalJSON() ([]byte, error) {
	//s := []rune(date.Format("2006-01-02 15:04:05"))
	//if date.Hour() < 10 {
	//	s = delChar(s, len(s)-8)
	//}
	return []byte(`"` + date.Format("2006-01-02 15:04:05") + `"`), nil
}

// All this to rune conversion is just to delete trailing zero in hour
// Golang dont have specified format tor this
func (date *DateTime) String() string {
	//s := []rune(date.Format("2006-01-02 15:04:05"))
	//if date.Hour() < 10 {
	//	s = delChar(s, len(s)-8)
	//}
	return date.Format("2006-01-02 15:04:05")
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}

func (date DateTime) Value() (driver.Value, error) {
	//date.Format("2006-01-02 15:04:05")
	return date.Time, nil
}

func (date *DateTime) Scan(src interface{}) (err error) {
	date.Time = src.(time.Time)
	return nil
}

type Decimal struct {
	decimal.Decimal
}

func (d *Decimal) String() string {
	return d.StringFixed(2)
}

func (d *Decimal) MarshalCSV() (string, error) {
	return d.StringFixed(2), nil
}

func (d *Decimal) UnmarshalCSV(csv string) (err error) {
	d.Decimal, err = decimal.NewFromString(csv)
	return err
}
