// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type TransactionStatus string

// Enum values for TransactionStatus
const (
	TransactionStatusAccepted TransactionStatus = "accepted"
	TransactionStatusDeclined TransactionStatus = "declined"
)

func AllTransactionStatus() []TransactionStatus {
	return []TransactionStatus{
		TransactionStatusAccepted,
		TransactionStatusDeclined,
	}
}

func (e TransactionStatus) IsValid() error {
	switch e {
	case TransactionStatusAccepted, TransactionStatusDeclined:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e TransactionStatus) String() string {
	return string(e)
}

type PaymentType string

// Enum values for PaymentType
const (
	PaymentTypeCash PaymentType = "cash"
	PaymentTypeCard PaymentType = "card"
)

func AllPaymentType() []PaymentType {
	return []PaymentType{
		PaymentTypeCash,
		PaymentTypeCard,
	}
}

func (e PaymentType) IsValid() error {
	switch e {
	case PaymentTypeCash, PaymentTypeCard:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e PaymentType) String() string {
	return string(e)
}
