// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/ledgers"
)

// LedgerEntry represents a row from 'ledger_entries'.
type LedgerEntry struct {
	ID              int32                  `json:"id"`                // id
	EntryID         string                 `json:"entry_id"`          // entry_id
	AccountID       string                 `json:"account_id"`        // account_id
	BranchID        string                 `json:"branch_id"`         // branch_id
	UserID          string                 `json:"user_id"`           // user_id
	TransactionID   string                 `json:"transaction_id"`    // transaction_id
	Type            LedgersEntryType       `json:"type"`              // type
	ProductID       string                 `json:"product_id"`        // product_id
	ProductType     LedgersProductType     `json:"product_type"`      // product_type
	Amount          []byte                 `json:"amount"`            // amount
	AccountingRate  CurrencyAccountingRate `json:"accounting_rate"`   // accounting_rate
	Notes           string                 `json:"notes"`             // notes
	ReversalEntryID string                 `json:"reversal_entry_id"` // reversal_entry_id
	CreateTime      time.Time              `json:"create_time"`       // create_time
	BookTime        time.Time              `json:"book_time"`         // book_time
	// xo fields
	Exists, Deleted bool
}

func NewLedgerEntry(pb *ledgers.Entry) (LedgerEntry, error) {
	if pb == nil {
		return LedgerEntry{}, ErrNilType{"LedgerEntry"}
	}
	le := LedgerEntry{
		EntryID:         pb.EntryID,
		AccountID:       pb.AccountID,
		BranchID:        pb.BranchID,
		UserID:          pb.UserID,
		TransactionID:   pb.TransactionID,
		Type:            NewLedgersEntryType(pb.Type),
		ProductID:       pb.ProductID,
		ProductType:     NewLedgersProductType(pb.ProductType),
		Notes:           pb.Notes,
		ReversalEntryID: pb.ReversalEntryID,
		CreateTime:      pb.CreateTime.AsTime(),
		BookTime:        pb.BookTime.AsTime(),
	}
	var err error
	le.Amount, err = protojson.Marshal(pb.Amount)
	if err != nil {
		return LedgerEntry{}, err
	}
	le.AccountingRate, err = NewCurrencyAccountingRate(pb.AccountingRate)
	if err != nil {
		return LedgerEntry{}, err
	}
	return le, nil
}

func (le LedgerEntry) PB() (*ledgers.Entry, error) {
	pb := &ledgers.Entry{
		EntryID:         le.EntryID,
		AccountID:       le.AccountID,
		BranchID:        le.BranchID,
		UserID:          le.UserID,
		TransactionID:   le.TransactionID,
		Type:            le.Type.PB(),
		ProductID:       le.ProductID,
		ProductType:     le.ProductType.PB(),
		Notes:           le.Notes,
		ReversalEntryID: le.ReversalEntryID,
		CreateTime:      timestamppb.New(le.CreateTime),
		BookTime:        timestamppb.New(le.BookTime),
	}
	var err error
	err = unmarshalMessage(le.Amount, &pb.Amount)
	if err != nil {
		return nil, err
	}
	pb.AccountingRate, err = le.AccountingRate.PB()
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type LedgerEntryRepository interface {
	InsertLedgerEntry(context.Context, *LedgerEntry) error
	ListLedgerEntries(context.Context, string, int32, string, *ListPosition) (ListStat, []*LedgerEntry, *ListPosition, error)

	// From ledger_entries_pkey
	LedgerEntryByID(context.Context, int32) (*LedgerEntry, error)

	UpdateLedgerEntryByID(context.Context, *LedgerEntry) error
	DeleteLedgerEntryByID(context.Context, int32) error

	// From ledger_entries_entry_id_idx
	LedgerEntryByEntryID(context.Context, string) (*LedgerEntry, error)

	UpdateLedgerEntryByEntryID(context.Context, *LedgerEntry) error
	DeleteLedgerEntryByEntryID(context.Context, string) error
}
