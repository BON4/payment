// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTransactonHistories(t *testing.T) {
	t.Parallel()

	query := TransactonHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTransactonHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactonHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TransactonHistories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactonHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactonHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactonHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TransactonHistoryExists(ctx, tx, o.TransactionId)
	if err != nil {
		t.Errorf("Unable to check if TransactonHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TransactonHistoryExists to return true, but got false.")
	}
}

func testTransactonHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	transactonHistoryFound, err := FindTransactonHistory(ctx, tx, o.TransactionId)
	if err != nil {
		t.Error(err)
	}

	if transactonHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTransactonHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TransactonHistories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTransactonHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TransactonHistories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTransactonHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactonHistoryOne := &TransactonHistory{}
	transactonHistoryTwo := &TransactonHistory{}
	if err = randomize.Struct(seed, transactonHistoryOne, transactonHistoryDBTypes, false, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, transactonHistoryTwo, transactonHistoryDBTypes, false, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactonHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactonHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TransactonHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTransactonHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	transactonHistoryOne := &TransactonHistory{}
	transactonHistoryTwo := &TransactonHistory{}
	if err = randomize.Struct(seed, transactonHistoryOne, transactonHistoryDBTypes, false, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, transactonHistoryTwo, transactonHistoryDBTypes, false, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactonHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactonHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func transactonHistoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func transactonHistoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactonHistory) error {
	*o = TransactonHistory{}
	return nil
}

func testTransactonHistoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TransactonHistory{}
	o := &TransactonHistory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TransactonHistory object: %s", err)
	}

	AddTransactonHistoryHook(boil.BeforeInsertHook, transactonHistoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	transactonHistoryBeforeInsertHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.AfterInsertHook, transactonHistoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	transactonHistoryAfterInsertHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.AfterSelectHook, transactonHistoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	transactonHistoryAfterSelectHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.BeforeUpdateHook, transactonHistoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	transactonHistoryBeforeUpdateHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.AfterUpdateHook, transactonHistoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	transactonHistoryAfterUpdateHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.BeforeDeleteHook, transactonHistoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	transactonHistoryBeforeDeleteHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.AfterDeleteHook, transactonHistoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	transactonHistoryAfterDeleteHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.BeforeUpsertHook, transactonHistoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	transactonHistoryBeforeUpsertHooks = []TransactonHistoryHook{}

	AddTransactonHistoryHook(boil.AfterUpsertHook, transactonHistoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	transactonHistoryAfterUpsertHooks = []TransactonHistoryHook{}
}

func testTransactonHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactonHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(transactonHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactonHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTransactonHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactonHistorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTransactonHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TransactonHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	transactonHistoryDBTypes = map[string]string{`TransactionId`: `bigint`, `RequestId`: `bigint`, `TerminalId`: `bigint`, `PartnerObjectId`: `bigint`, `AmountTotal`: `bigint`, `AmountOriginal`: `bigint`, `CommissionPS`: `numeric`, `CommissionClient`: `numeric`, `CommissionProvider`: `numeric`, `DateInput`: `timestamp without time zone`, `DatePost`: `timestamp without time zone`, `Status`: `enum.transaction_status('accepted','declined')`, `PaymentType`: `enum.payment_type('cash','card')`, `PaymentNumber`: `text`, `ServiceId`: `bigint`, `Service`: `text`, `PayeeId`: `bigint`, `PayeeName`: `text`, `PayeeBankMfo`: `bigint`, `PayeeBankAccount`: `text`, `PaymentNarrative`: `text`}
	_                        = bytes.MinRead
)

func testTransactonHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(transactonHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(transactonHistoryAllColumns) == len(transactonHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTransactonHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(transactonHistoryAllColumns) == len(transactonHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TransactonHistory{}
	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactonHistoryDBTypes, true, transactonHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(transactonHistoryAllColumns, transactonHistoryPrimaryKeyColumns) {
		fields = transactonHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			transactonHistoryAllColumns,
			transactonHistoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TransactonHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTransactonHistoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(transactonHistoryAllColumns) == len(transactonHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TransactonHistory{}
	if err = randomize.Struct(seed, &o, transactonHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TransactonHistory: %s", err)
	}

	count, err := TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, transactonHistoryDBTypes, false, transactonHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactonHistory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TransactonHistory: %s", err)
	}

	count, err = TransactonHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}