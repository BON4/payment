package domain

import (
	"context"
	"time"
)

const TIME_FORMAT = "2022-08-23 11:25:27"

type FindTxRequest struct {
	TransactionID *int64  `form:"transaction_id" json:"transaction_id,omitempty"`
	TerminalIDs   []int64 `form:"terminal_id" json:"terminal_id,omitempty"`
	Status        *string `form:"status" json:"status,omitempty"`
	PaymentType   *string `form:"payment_type" json:"payment_type,omitempty"`

	PaymentNarrative *string `form:"payment_narrative" json:"payment_narrative,omitempty"`

	PostDateFrom *time.Time `form:"post_date_from" json:"post_date_from,omitempty"`
	PostDateTo   *time.Time `form:"post_date_to" json:"post_date_to,omitempty"`

	PageSize   int64 `form:"bpage_size" json:"page_size" binding:"required"`
	PageNumber int64 `form:"page_number" json:"page_number" binding:"required"`
}

type TxUsecase interface {
	Create(ctx context.Context, tx *TransactonHistory) error
	List(ctx context.Context, cond FindTxRequest) ([]*TransactonHistory, error)
}
