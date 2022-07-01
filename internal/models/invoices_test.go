// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testInvoices(t *testing.T) {
	t.Parallel()

	query := Invoices()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testInvoicesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
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

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoicesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Invoices().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoicesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InvoiceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoicesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := InvoiceExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Invoice exists: %s", err)
	}
	if !e {
		t.Errorf("Expected InvoiceExists to return true, but got false.")
	}
}

func testInvoicesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	invoiceFound, err := FindInvoice(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if invoiceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testInvoicesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Invoices().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testInvoicesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Invoices().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testInvoicesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	invoiceOne := &Invoice{}
	invoiceTwo := &Invoice{}
	if err = randomize.Struct(seed, invoiceOne, invoiceDBTypes, false, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}
	if err = randomize.Struct(seed, invoiceTwo, invoiceDBTypes, false, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = invoiceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = invoiceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Invoices().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testInvoicesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	invoiceOne := &Invoice{}
	invoiceTwo := &Invoice{}
	if err = randomize.Struct(seed, invoiceOne, invoiceDBTypes, false, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}
	if err = randomize.Struct(seed, invoiceTwo, invoiceDBTypes, false, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = invoiceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = invoiceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testInvoicesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInvoicesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(invoiceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInvoicesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
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

func testInvoicesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InvoiceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testInvoicesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Invoices().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	invoiceDBTypes = map[string]string{`ID`: `uuid`, `UserID`: `uuid`, `Title`: `character varying`, `Network`: `character varying`, `Recipient`: `character varying`, `Amount`: `bigint`, `JumpURL`: `character varying`, `NotifyURL`: `character varying`, `Status`: `character varying`, `Metadata`: `character varying`, `TXNHash`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testInvoicesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(invoicePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(invoiceAllColumns) == len(invoicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testInvoicesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(invoiceAllColumns) == len(invoicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Invoice{}
	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, invoiceDBTypes, true, invoicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(invoiceAllColumns, invoicePrimaryKeyColumns) {
		fields = invoiceAllColumns
	} else {
		fields = strmangle.SetComplement(
			invoiceAllColumns,
			invoicePrimaryKeyColumns,
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

	slice := InvoiceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testInvoicesUpsert(t *testing.T) {
	t.Parallel()

	if len(invoiceAllColumns) == len(invoicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Invoice{}
	if err = randomize.Struct(seed, &o, invoiceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Invoice: %s", err)
	}

	count, err := Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, invoiceDBTypes, false, invoicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Invoice: %s", err)
	}

	count, err = Invoices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}