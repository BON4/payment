package usecase

import (
	"context"
	"database/sql"
	"os"

	"github.com/BON4/payment/internal/domain"
	boilmodels "github.com/BON4/payment/internal/domain/boil_postgres"
	"github.com/BON4/payment/internal/pkg/csvupload"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type txUsecase struct {
	db  *sql.DB
	csv *csvupload.CSVUploader[domain.TransactonHistory]
}

func NewTxUsecase(db *sql.DB) domain.TxUsecase {
	return &txUsecase{
		db:  db,
		csv: csvupload.NewCSVUploader[domain.TransactonHistory](db),
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

func (t *txUsecase) CSVCopyInsert(ctx context.Context, filePath string) (int64, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	par, err := domain.NewTxHitoryParser(file)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	return t.csv.Upload(context.Background(), "boil", par)
}

func (t *txUsecase) List(ctx context.Context, cond domain.FindTxRequest) ([]*domain.TransactonHistory, error) {
	var conds []qm.QueryMod = make([]qm.QueryMod, 0, 1)

	if cond.TransactionID != nil {
		conds = append(conds, qm.Where(`"TransactionId"=?`, cond.TransactionID))
	}

	if len(cond.TerminalIDs) > 0 {
		//TODO: not my problem or sqlboiler problem. Rather the Go itself.
		//Maby try sqlx.In()
		convertedIDs := make([]interface{}, len(cond.TerminalIDs))
		for index, num := range cond.TerminalIDs {
			convertedIDs[index] = num
		}
		conds = append(conds, qm.WhereIn(`"TerminalId" IN ?`, convertedIDs...))
	}

	if cond.Status != nil {
		conds = append(conds, qm.Where(`"Status"=?`, cond.Status))
	}

	if cond.PaymentType != nil {
		conds = append(conds, qm.Where(`"PaymentType"=?`, cond.PaymentType))
	}

	if cond.PaymentNarrative != nil {
		conds = append(conds, qm.Where(`"PaymentNarrative" LIKE ?`, "%"+*cond.PaymentNarrative+"%"))
	}

	if cond.PostDateFrom != nil {
		conds = append(conds, qm.Where(`"DatePost" >= ?`, cond.PostDateFrom))
	}

	if cond.PostDateTo != nil {
		conds = append(conds, qm.Where(`"DatePost" < ?`, cond.PostDateTo.AddDate(0, 0, 1)))
	}

	if cond.PageSize > 0 {
		conds = append(conds, qm.Offset(int(cond.PageNumber*cond.PageSize)), qm.Limit(int(cond.PageSize)))
	}

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
