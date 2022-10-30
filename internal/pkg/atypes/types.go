package atypes

import (
	"database/sql/driver"
	"time"

	"fmt"

	"github.com/shopspring/decimal"
)

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

type DateTime struct {
	time.Time
}

// Golang dont have specified format tor h-mm-ss
func (date *DateTime) MarshalCSV() (string, error) {
	return fmt.Sprintf("%d-%02d-%02d %01d:%02d:%02d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second()), nil
}

func (date *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d-%02d-%02d %01d:%02d:%02d"`,
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())), nil
}

// Golang dont have specified format tor h-mm-ss
func (date *DateTime) String() string {
	return fmt.Sprintf("%d-%02d-%02d %01d:%02d:%02d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())
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
