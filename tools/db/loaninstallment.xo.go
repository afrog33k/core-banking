// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/loans"
)

// LoanInstallment represents a row from 'loan_installments'.
type LoanInstallment struct {
	ID               int32                 `json:"id"`                 // id
	InstallmentID    string                `json:"installment_id"`     // installment_id
	Fees             []byte                `json:"fees"`               // fees
	Interest         []byte                `json:"interest"`           // interest
	IsPaymentHoliday bool                  `json:"is_payment_holiday"` // is_payment_holiday
	PayTime          sql.NullTime          `json:"pay_time"`           // pay_time
	Order            string                `json:"order"`              // order
	ParentAccountID  string                `json:"parent_account_id"`  // parent_account_id
	Penalty          []byte                `json:"penalty"`            // penalty
	Principal        []byte                `json:"principal"`          // principal
	DueTime          time.Time             `json:"due_time"`           // due_time
	RepayTime        time.Time             `json:"repay_time"`         // repay_time
	State            LoansInstallmentState `json:"state"`              // state
	// xo fields
	Exists, Deleted bool
}

func NewLoanInstallment(pb *loans.Installment) (LoanInstallment, error) {
	if pb == nil {
		return LoanInstallment{}, ErrNilType{"LoanInstallment"}
	}
	li := LoanInstallment{
		InstallmentID:    pb.InstallmentID,
		IsPaymentHoliday: pb.IsPaymentHoliday,
		PayTime:          toNullTime(pb.PayTime),
		Order:            pb.Order,
		ParentAccountID:  pb.ParentAccountID,
		DueTime:          pb.DueTime.AsTime(),
		RepayTime:        pb.RepayTime.AsTime(),
		State:            NewLoansInstallmentState(pb.State),
	}
	var err error
	li.Fees, err = marshalArray(pb.Fees)
	if err != nil {
		return LoanInstallment{}, err
	}
	li.Interest, err = protojson.Marshal(pb.Interest)
	if err != nil {
		return LoanInstallment{}, err
	}
	li.Penalty, err = protojson.Marshal(pb.Penalty)
	if err != nil {
		return LoanInstallment{}, err
	}
	li.Principal, err = protojson.Marshal(pb.Principal)
	if err != nil {
		return LoanInstallment{}, err
	}
	return li, nil
}

func (li LoanInstallment) PB() (*loans.Installment, error) {
	pb := &loans.Installment{
		InstallmentID:    li.InstallmentID,
		IsPaymentHoliday: li.IsPaymentHoliday,
		PayTime:          toTimePB(li.PayTime),
		Order:            li.Order,
		ParentAccountID:  li.ParentAccountID,
		DueTime:          timestamppb.New(li.DueTime),
		RepayTime:        timestamppb.New(li.RepayTime),
		State:            li.State.PB(),
	}
	var err error
	err = unmarshalArray(li.Fees, &pb.Fees)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(li.Interest, &pb.Interest)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(li.Penalty, &pb.Penalty)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(li.Principal, &pb.Principal)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type LoanInstallmentRepository interface {
	InsertLoanInstallment(context.Context, *LoanInstallment) error
	ListLoanInstallments(context.Context, string, int32, string, *ListPosition) (ListStat, []*LoanInstallment, *ListPosition, error)

	// From loan_installments_pkey
	LoanInstallmentByID(context.Context, int32) (*LoanInstallment, error)

	UpdateLoanInstallmentByID(context.Context, *LoanInstallment) error
	DeleteLoanInstallmentByID(context.Context, int32) error

	// From loan_installments_installment_id_idx
	LoanInstallmentByInstallmentID(context.Context, string) (*LoanInstallment, error)

	UpdateLoanInstallmentByInstallmentID(context.Context, *LoanInstallment) error
	DeleteLoanInstallmentByInstallmentID(context.Context, string) error
}
