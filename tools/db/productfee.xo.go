// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/products"
)

// ProductFee represents a row from 'product_fees'.
type ProductFee struct {
	ID                   int32                        `json:"id"`                    // id
	FeeID                string                       `json:"fee_id"`                // fee_id
	Name                 string                       `json:"name"`                  // name
	CalculationMethod    ProductsFeeCalculationMethod `json:"calculation_method"`    // calculation_method
	Required             bool                         `json:"required"`              // required
	Amount               []byte                       `json:"amount"`                // amount
	Percent              string                       `json:"percent"`               // percent
	ApplyDateMethod      ProductsFeeApplyDateMethod   `json:"apply_date_method"`     // apply_date_method
	Trigger              ProductsFeeTrigger           `json:"trigger"`               // trigger
	AccountingRules      []byte                       `json:"accounting_rules"`      // accounting_rules
	AmortizationSettings []byte                       `json:"amortization_settings"` // amortization_settings
	IsActive             bool                         `json:"is_active"`             // is_active
	IsTaxable            bool                         `json:"is_taxable"`            // is_taxable
	CreateTime           time.Time                    `json:"create_time"`           // create_time
	UpdateTime           time.Time                    `json:"update_time"`           // update_time
	// xo fields
	Exists, Deleted bool
}

func NewProductFee(pb *products.Fee) (ProductFee, error) {
	if pb == nil {
		return ProductFee{}, ErrNilType{"ProductFee"}
	}
	pf := ProductFee{
		FeeID:             pb.FeeID,
		Name:              pb.Name,
		CalculationMethod: NewProductsFeeCalculationMethod(pb.CalculationMethod),
		Required:          pb.Required,
		Percent:           pb.Percent,
		ApplyDateMethod:   NewProductsFeeApplyDateMethod(pb.ApplyDateMethod),
		Trigger:           NewProductsFeeTrigger(pb.Trigger),
		IsActive:          pb.IsActive,
		IsTaxable:         pb.IsTaxable,
		CreateTime:        pb.CreateTime.AsTime(),
		UpdateTime:        pb.UpdateTime.AsTime(),
	}
	var err error
	pf.Amount, err = protojson.Marshal(pb.Amount)
	if err != nil {
		return ProductFee{}, err
	}
	pf.AccountingRules, err = marshalArray(pb.AccountingRules)
	if err != nil {
		return ProductFee{}, err
	}
	pf.AmortizationSettings, err = protojson.Marshal(pb.AmortizationSettings)
	if err != nil {
		return ProductFee{}, err
	}
	return pf, nil
}

func (pf ProductFee) PB() (*products.Fee, error) {
	pb := &products.Fee{
		FeeID:             pf.FeeID,
		Name:              pf.Name,
		CalculationMethod: pf.CalculationMethod.PB(),
		Required:          pf.Required,
		Percent:           pf.Percent,
		ApplyDateMethod:   pf.ApplyDateMethod.PB(),
		Trigger:           pf.Trigger.PB(),
		IsActive:          pf.IsActive,
		IsTaxable:         pf.IsTaxable,
		CreateTime:        timestamppb.New(pf.CreateTime),
		UpdateTime:        timestamppb.New(pf.UpdateTime),
	}
	var err error
	err = unmarshalMessage(pf.Amount, &pb.Amount)
	if err != nil {
		return nil, err
	}
	err = unmarshalArray(pf.AccountingRules, &pb.AccountingRules)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(pf.AmortizationSettings, &pb.AmortizationSettings)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type ProductFeeRepository interface {
	InsertProductFee(context.Context, *ProductFee) error
	ListProductFees(context.Context, string, int32, string, *ListPosition) (ListStat, []*ProductFee, *ListPosition, error)

	// From product_fees_pkey
	ProductFeeByID(context.Context, int32) (*ProductFee, error)

	UpdateProductFeeByID(context.Context, *ProductFee) error
	DeleteProductFeeByID(context.Context, int32) error

	// From product_fees_fee_id_idx
	ProductFeeByFeeID(context.Context, string) (*ProductFee, error)

	UpdateProductFeeByFeeID(context.Context, *ProductFee) error
	DeleteProductFeeByFeeID(context.Context, string) error
}
