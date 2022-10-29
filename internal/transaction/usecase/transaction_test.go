package usecase_test

import (
	"bytes"
	"context"
	"database/sql"
	"io"
	"os"
	"testing"
	"time"

	"github.com/BON4/payment/internal/domain"
	"github.com/BON4/payment/internal/pkg/dbtest"
	"github.com/BON4/payment/internal/transaction/usecase"
	_ "github.com/lib/pq"
)

var db *sql.DB
var testExampleCSV = "example.csv"
var connStr = "postgresql://root:secret@localhost:5432/payment_test?sslmode=disable"

func lineCounter(r io.Reader) (int64, error) {
	buf := make([]byte, 32*1024)
	count := int64(0)
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += int64(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func TestMain(m *testing.M) {
	var err error
	db, err = dbtest.ConnectTestDB("postgres", connStr)

	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestCSVCopyInsert(t *testing.T) {
	defer func() {
		err := dbtest.FlushDB(t, db, "transacton_history")
		if err != nil {
			t.Fatal(err)
		}
	}()

	file, err := os.OpenFile(testExampleCSV, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	m, err := lineCounter(file)
	if err != nil {
		t.Fatal(err)
		return
	}

	file.Seek(0, 0)

	tuc := usecase.NewTxUsecase(db)

	n, err := tuc.CSVInsert(context.Background(), file)
	if err != nil {
		t.Fatal(err)
		return
	}

	if n != m {
		t.Fatalf("Number of inserted csv objects not equal to original number of objects in csv file: %d != %d", n, m)
		return
	}
}

func TestList(t *testing.T) {
	defer func() {
		err := dbtest.FlushDB(t, db, "transacton_history")
		if err != nil {
			t.Fatal(err)
		}
	}()

	file, err := os.OpenFile(testExampleCSV, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	tuc := usecase.NewTxUsecase(db)

	n, err := tuc.CSVInsert(context.Background(), file)
	if err != nil {
		t.Fatal(err)
		return
	}

	txs, err := tuc.List(context.Background(), domain.FindTxRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	// for _, tx := range txs {
	// 	t.Logf("%+v\n", tx)
	// }

	if int64(len(txs)) != n {
		t.Fatalf("Number of inserted csv objects not equal to original number of objects in csv file: %d != %d", n, len(txs))
		return
	}

	t.Run("вивантаження попередньо збережених даних в JSON ok", func(t *testing.T) {
		cond := domain.FindTxRequest{
			PageSize:   10,
			PageNumber: 1,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if int64(len(txs)) != cond.PageSize {
			t.Fatalf("Number of inserted csv objects not equal to original number of objects in csv file: %d != %d", cond.PageSize, len(txs))
			return
		}

	})

	t.Run("пошук по transaction_id ok", func(t *testing.T) {
		id := int64(22)
		cond := domain.FindTxRequest{
			TransactionID: &id,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) != 1 {
			t.Fatalf("Number of inserted csv objects not equal to original number of objects in csv file: %d != %d", 1, len(txs))
			return
		}

	})

	t.Run("пошук по transaction_id fail", func(t *testing.T) {
		id := int64(2222)
		cond := domain.FindTxRequest{
			TransactionID: &id,
		}

		txs, err := tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) != 0 {
			t.Fatalf("No data shoud be returned")
			return
		}
	})

	t.Run("пошук по terminal_id (можливість вказати декілька одночасно id) ok", func(t *testing.T) {
		cond := domain.FindTxRequest{
			TerminalIDs: []int64{3511, 3570, 3542},
		}

		txs, err := tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) != 3 {
			t.Fatalf("Not enough objects returned")
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shuld be returned")
			return
		}
	})

	t.Run("пошук по terminal_id (можливість вказати декілька одночасно id) fail", func(t *testing.T) {
		cond := domain.FindTxRequest{
			TerminalIDs: []int64{693511, 1000020, 4201337},
		}

		txs, err := tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) != 0 {
			t.Fatalf("No data shoud be returned")
			return
		}
	})

	t.Run("пошук по status (accepted/declined) ok", func(t *testing.T) {
		status := "accepted"
		cond := domain.FindTxRequest{
			Status: &status,
		}

		txs, err := tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}

		status = "declined"
		cond = domain.FindTxRequest{
			Status: &status,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}
	})

	t.Run("пошук по status (accepted/declined) fail", func(t *testing.T) {
		status := "err"
		cond := domain.FindTxRequest{
			Status: &status,
		}

		_, err := tuc.List(context.Background(), cond)
		if err == nil {
			t.Fatalf("Shoud be error invalid input value for enum")
			return
		}
	})

	t.Run("пошук по payment_type (cash/card) ok", func(t *testing.T) {
		status := "cash"
		cond := domain.FindTxRequest{
			PaymentType: &status,
		}

		txs, err := tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}

		status = "card"
		cond = domain.FindTxRequest{
			PaymentType: &status,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}
	})

	t.Run("пошук по payment_type (cash/card) fail", func(t *testing.T) {
		status := "err"
		cond := domain.FindTxRequest{
			PaymentType: &status,
		}

		_, err := tuc.List(context.Background(), cond)
		if err == nil {
			t.Fatalf("Shoud be error invalid input value for enum")
			return
		}
	})

	t.Run("пошук по date_post по періодам (from/to) ok", func(t *testing.T) {
		from, err := time.Parse("2006-01-02", "2022-08-17")
		if err != nil {
			t.Fatal(err)
			return
		}

		cond := domain.FindTxRequest{
			PostDateFrom: &from,
		}

		txs, err := tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}

		to, err := time.Parse("2006-01-02", "2022-08-17")
		if err != nil {
			t.Fatal(err)
			return
		}

		cond = domain.FindTxRequest{
			PostDateTo: &to,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}

		to = to.AddDate(0, 0, 6)
		cond = domain.FindTxRequest{
			PostDateFrom: &from,
			PostDateTo:   &to,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}
	})

	t.Run("пошук по date_post по періодам (from/to) ok", func(t *testing.T) {
		from, err := time.Parse("2006-01-02", "2022-08-17")
		if err != nil {
			t.Fatal(err)
			return
		}

		cond := domain.FindTxRequest{
			PostDateFrom: &from,
		}

		txs, err := tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}

		to, err := time.Parse("2006-01-02", "2022-08-17")
		if err != nil {
			t.Fatal(err)
			return
		}

		cond = domain.FindTxRequest{
			PostDateTo: &to,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}

		to = to.AddDate(0, 0, 6)
		cond = domain.FindTxRequest{
			PostDateFrom: &from,
			PostDateTo:   &to,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}
	})

	t.Run("пошук по date_post по періодам (from/to) fail", func(t *testing.T) {
		from, err := time.Parse("2006-01-02", "2022-08-17")
		if err != nil {
			t.Fatal(err)
			return
		}

		to, err := time.Parse("2006-01-02", "2022-04-17")
		if err != nil {
			t.Fatal(err)
			return
		}

		cond := domain.FindTxRequest{
			PostDateFrom: &from,
			PostDateTo:   &to,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) != 0 {
			t.Fatalf("No data shoud be returned")
			return
		}
	})

	t.Run("пошук по частково вказаному payment_narrative ok", func(t *testing.T) {
		pNarr := "122"
		cond := domain.FindTxRequest{
			PaymentNarrative: &pNarr,
		}

		txs, err = tuc.List(context.Background(), cond)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(txs) == 0 {
			t.Fatalf("Data shoud be returned")
			return
		}
	})
	// for _, tx := range txs {
	// 	t.Logf("%+v\n", tx)
	// }

}
