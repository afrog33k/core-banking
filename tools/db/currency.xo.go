// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"bnk.to/core/api/v1/currencies"
)

// Currency represents a row from 'currencies'.
type Currency struct {
	ID       int32  `json:"id"`       // id
	Code     string `json:"code"`     // code
	Name     string `json:"name"`     // name
	Symbol   string `json:"symbol"`   // symbol
	Format   string `json:"format"`   // format
	Decimals int32  `json:"decimals"` // decimals
	// xo fields
	Exists, Deleted bool
}

func NewCurrency(pb *currencies.Currency) (Currency, error) {
	if pb == nil {
		return Currency{}, ErrNilType{"Currency"}
	}
	c := Currency{
		Code:     pb.Code,
		Name:     pb.Name,
		Symbol:   pb.Symbol,
		Format:   pb.Format,
		Decimals: pb.Decimals,
	}
	return c, nil
}

func (c Currency) PB() (*currencies.Currency, error) {
	pb := &currencies.Currency{
		Code:     c.Code,
		Name:     c.Name,
		Symbol:   c.Symbol,
		Format:   c.Format,
		Decimals: c.Decimals,
	}
	return pb, nil
}

type CurrencyRepository interface {
	InsertCurrency(context.Context, *Currency) error
	ListCurrencies(context.Context, string, int32, string, *ListPosition) (ListStat, []*Currency, *ListPosition, error)

	// From currencies_pkey
	CurrencyByID(context.Context, int32) (*Currency, error)

	UpdateCurrencyByID(context.Context, *Currency) error
	DeleteCurrencyByID(context.Context, int32) error

	// From currencies_code_idx
	CurrencyByCode(context.Context, string) (*Currency, error)

	UpdateCurrencyByCode(context.Context, *Currency) error
	DeleteCurrencyByCode(context.Context, string) error

	// From currencies_name_idx
	CurrencyByName(context.Context, string) (*Currency, error)

	UpdateCurrencyByName(context.Context, *Currency) error
	DeleteCurrencyByName(context.Context, string) error
}
