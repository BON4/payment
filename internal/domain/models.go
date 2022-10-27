package domain

import (
	boilmodels "github.com/BON4/payment/internal/domain/boil_postgres"
	"github.com/BON4/payment/internal/pkg/atypes"
)

type TransactionStatus string

// Enum values for TransactionStatus
const (
	TransactionStatusAccepted TransactionStatus = "accepted"
	TransactionStatusDeclined TransactionStatus = "declined"
)

type PaymentType string

// Enum values for PaymentType
const (
	PaymentTypeCash PaymentType = "cash"
	PaymentTypeCard PaymentType = "card"
)

type TransactonHistory struct {
	TransactionId      int64             `csv:"TransactionId" boil:"TransactionId" json:"TransactionId" toml:"TransactionId" yaml:"TransactionId"`
	RequestId          int64             `csv:"RequestId" boil:"RequestId" json:"RequestId" toml:"RequestId" yaml:"RequestId"`
	TerminalId         int64             `csv:"TerminalId" boil:"TerminalId" json:"TerminalId" toml:"TerminalId" yaml:"TerminalId"`
	PartnerObjectId    int64             `csv:"PartnerObjectId" boil:"PartnerObjectId" json:"PartnerObjectId" toml:"PartnerObjectId" yaml:"PartnerObjectId"`
	AmountTotal        int64             `csv:"AmountTotal" boil:"AmountTotal" json:"AmountTotal" toml:"AmountTotal" yaml:"AmountTotal"`
	AmountOriginal     int64             `csv:"AmountOriginal" boil:"AmountOriginal" json:"AmountOriginal" toml:"AmountOriginal" yaml:"AmountOriginal"`
	CommissionPS       atypes.Decimal    `csv:"CommissionPS" boil:"CommissionPS" json:"CommissionPS" toml:"CommissionPS" yaml:"CommissionPS"`
	CommissionClient   atypes.Decimal    `csv:"CommissionClient" boil:"CommissionClient" json:"CommissionClient" toml:"CommissionClient" yaml:"CommissionClient"`
	CommissionProvider atypes.Decimal    `csv:"CommissionProvider" boil:"CommissionProvider" json:"CommissionProvider" toml:"CommissionProvider" yaml:"CommissionProvider"`
	DateInput          atypes.DateTime   `csv:"DateInput" boil:"DateInput" json:"DateInput" toml:"DateInput" yaml:"DateInput"`
	DatePost           atypes.DateTime   `csv:"DatePost" boil:"DatePost" json:"DatePost" toml:"DatePost" yaml:"DatePost"`
	Status             TransactionStatus `csv:"Status" boil:"Status" json:"Status" toml:"Status" yaml:"Status"`
	PaymentType        PaymentType       `csv:"PaymentType" boil:"PaymentType" json:"PaymentType" toml:"PaymentType" yaml:"PaymentType"`
	PaymentNumber      string            `csv:"PaymentNumber" boil:"PaymentNumber" json:"PaymentNumber" toml:"PaymentNumber" yaml:"PaymentNumber"`
	ServiceId          int64             `csv:"ServiceId" boil:"ServiceId" json:"ServiceId" toml:"ServiceId" yaml:"ServiceId"`
	Service            string            `csv:"Service" boil:"Service" json:"Service" toml:"Service" yaml:"Service"`
	PayeeId            int64             `csv:"PayeeId" boil:"PayeeId" json:"PayeeId" toml:"PayeeId" yaml:"PayeeId"`
	PayeeName          string            `csv:"PayeeName" boil:"PayeeName" json:"PayeeName" toml:"PayeeName" yaml:"PayeeName"`
	PayeeBankMfo       int64             `csv:"PayeeBankMfo" boil:"PayeeBankMfo" json:"PayeeBankMfo" toml:"PayeeBankMfo" yaml:"PayeeBankMfo"`
	PayeeBankAccount   string            `csv:"PayeeBankAccount" boil:"PayeeBankAccount" json:"PayeeBankAccount" toml:"PayeeBankAccount" yaml:"PayeeBankAccount"`
	PaymentNarrative   string            `csv:"PaymentNarrative" boil:"PaymentNarrative" json:"PaymentNarrative" toml:"PaymentNarrative" yaml:"PaymentNarrative"`
}

func BoilToDomainBinding(in *boilmodels.TransactonHistory, out *TransactonHistory) {
	if in != nil && out != nil {
		out.TransactionId = in.TransactionId
		out.RequestId = in.RequestId
		out.TerminalId = in.TerminalId
		out.PartnerObjectId = in.PartnerObjectId
		out.AmountTotal = in.AmountTotal
		out.AmountOriginal = in.AmountOriginal
		out.CommissionPS = in.CommissionPS
		out.CommissionClient = in.CommissionClient
		out.CommissionProvider = in.CommissionProvider
		out.DateInput = in.DateInput
		out.DatePost = in.DatePost
		out.Status = TransactionStatus(in.Status)
		out.PaymentType = PaymentType(in.PaymentType)
		out.PaymentNumber = in.PaymentNumber
		out.ServiceId = in.ServiceId
		out.Service = in.Service
		out.PayeeId = in.PayeeId
		out.PayeeName = in.PayeeName
		out.PayeeBankMfo = in.PayeeBankMfo
		out.PayeeBankAccount = in.PayeeBankAccount
		out.PaymentNarrative = in.PaymentNarrative
	}
}

func DomainToBoilBinding(in *TransactonHistory, out *boilmodels.TransactonHistory) {
	if in != nil && out != nil {

		out.TransactionId = in.TransactionId
		out.RequestId = in.RequestId
		out.TerminalId = in.TerminalId
		out.PartnerObjectId = in.PartnerObjectId
		out.AmountTotal = in.AmountTotal
		out.AmountOriginal = in.AmountOriginal
		out.CommissionPS = in.CommissionPS
		out.CommissionClient = in.CommissionClient
		out.CommissionProvider = in.CommissionProvider
		out.DateInput = in.DateInput
		out.DatePost = in.DatePost
		out.Status = boilmodels.TransactionStatus(in.Status)
		out.PaymentType = boilmodels.PaymentType(in.PaymentType)
		out.PaymentNumber = in.PaymentNumber
		out.ServiceId = in.ServiceId
		out.Service = in.Service
		out.PayeeId = in.PayeeId
		out.PayeeName = in.PayeeName
		out.PayeeBankMfo = in.PayeeBankMfo
		out.PayeeBankAccount = in.PayeeBankAccount
		out.PaymentNarrative = in.PaymentNarrative
	}
}
