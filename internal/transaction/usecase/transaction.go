package usecase

import (
	"context"
	"database/sql"

	"github.com/BON4/payment/internal/domain"
	boilmodels "github.com/BON4/payment/internal/domain/boil_postgres"
	"github.com/fatih/structs"
	"github.com/lib/pq"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type txUsecase struct {
	db *sql.DB
}

func NewTxUsecase(db *sql.DB) domain.TxUsecase {
	return &txUsecase{
		db: db,
	}
}

func (t *txUsecase) Create(ctx context.Context, txhist *domain.TransactonHistory) error {
	boilTx := &boilmodels.TransactonHistory{}
	domain.DomainToBoilBinding(txhist, boilTx)

	if err := boilTx.Insert(ctx, t.db, boil.Infer()); err != nil {
		return err
	}

	domain.BoilToDomainBinding(boilTx, txhist)

	return nil
}

func (t *txUsecase) BulkInsert(ctx context.Context, data []*domain.TransactonHistory) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	var boilTx boilmodels.TransactonHistory
	for _, d := range data {
		domain.DomainToBoilBinding(d, &boilTx)
		err := boilTx.Insert(ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (t *txUsecase) CSVCopy(ctx context.Context, filePath string) error {
	return nil
}

func (t *txUsecase) CopyInsert(ctx context.Context, data []*domain.TransactonHistory) error {

	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	names := make([]string, 0)

	for _, field := range structs.Fields(domain.TransactonHistory{}) {
		names = append(names, field.Tag("boil"))
	}

	stmt, err := tx.PrepareContext(ctx, pq.CopyInSchema("db", "table", names...))
	if err != nil {
		return err
	}

	for _, t := range data {

		_, err := stmt.Exec(t.Transactionid, t.Requestid, t.Terminalid, t.Partnerobjectid, t.Amounttotal, t.Amountoriginal, t.Commissionclient, t.Commissionprovider, t.Dateinput, t.Datepost, t.Status, t.Paymenttype, t.Paymentnumber, t.Serviceid, t.Service, t.Payeeid, t.Payeename, t.Payeebankmfo, t.Payeebankaccount, t.Paymentnarrative)
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	err = stmt.Close()
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (t *txUsecase) List(ctx context.Context, cond domain.FindTxRequest) ([]*domain.TransactonHistory, error) {
	var conds []qm.QueryMod = make([]qm.QueryMod, 0, 1)

	if cond.TransactionID != nil {
		conds = append(conds, qm.Where("transactionid=?", cond.TransactionID))
	}

	if len(cond.TerminalIDs) > 0 {
		conds = append(conds, qm.WhereIn("terminalid in ?", cond.TerminalIDs))
	}

	if cond.Status != nil {
		conds = append(conds, qm.Where("status=?", cond.Status))
	}

	if cond.PaymentType != nil {
		conds = append(conds, qm.Where("PaymentType=?", cond.PaymentType))
	}

	if cond.PaymentNarrative != nil {
		conds = append(conds, qm.Where("PaymentNarrative like %?%", cond.PaymentType))
	}

	if cond.PostDateFrom != nil {
		year, mouth, day := cond.PostDateFrom.Date()
		conds = append(conds, qm.Where("DatePost >= '?-?-?'::date", year, mouth, day))
	}

	if cond.PostDateTo != nil {
		year, mouth, day := cond.PostDateTo.Date()
		//TODO: check if +'1 day' is needed
		conds = append(conds, qm.Where("DatePost < '?-?-?'::date + '1 day'::date", year, mouth, day))
	}

	conds = append(conds, qm.Offset(int(cond.PageNumber)), qm.Limit(int(cond.PageSize)))

	txs, err := boilmodels.TransactonHistories(conds...).All(ctx, t.db)
	if err != nil {
		return []*domain.TransactonHistory{}, err
	}

	domainTxs := make([]*domain.TransactonHistory, len(txs))

	for i, tx := range txs {
		domainTxs[i] = &domain.TransactonHistory{}
		domain.BoilToDomainBinding(tx, domainTxs[i])
	}

	txs = nil

	return domainTxs, nil
}
