// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"

	"bnk.to/core/api/v1/loans"
)

// LoanInvestorFund represents a row from 'loan_investor_funds'.
type LoanInvestorFund struct {
	ID                 int32   `json:"id"`                  // id
	FundID             string  `json:"fund_id"`             // fund_id
	Amount             []byte  `json:"amount"`              // amount
	InterestCommission float64 `json:"interest_commission"` // interest_commission
	SharePercentage    float64 `json:"share_percentage"`    // share_percentage
	// xo fields
	Exists, Deleted bool
}

func NewLoanInvestorFund(pb *loans.InvestorFund) (LoanInvestorFund, error) {
	if pb == nil {
		return LoanInvestorFund{}, ErrNilType{"LoanInvestorFund"}
	}
	lif := LoanInvestorFund{
		FundID:             pb.FundID,
		InterestCommission: pb.InterestCommission,
		SharePercentage:    pb.SharePercentage,
	}
	var err error
	lif.Amount, err = protojson.Marshal(pb.Amount)
	if err != nil {
		return LoanInvestorFund{}, err
	}
	return lif, nil
}

func (lif LoanInvestorFund) PB() (*loans.InvestorFund, error) {
	pb := &loans.InvestorFund{
		FundID:             lif.FundID,
		InterestCommission: lif.InterestCommission,
		SharePercentage:    lif.SharePercentage,
	}
	var err error
	err = unmarshalMessage(lif.Amount, &pb.Amount)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type LoanInvestorFundRepository interface {
	InsertLoanInvestorFund(context.Context, *LoanInvestorFund) error
	ListLoanInvestorFunds(context.Context, string, int32, string, *ListPosition) (ListStat, []*LoanInvestorFund, *ListPosition, error)

	// From loan_investor_funds_pkey
	LoanInvestorFundByID(context.Context, int32) (*LoanInvestorFund, error)

	UpdateLoanInvestorFundByID(context.Context, *LoanInvestorFund) error
	DeleteLoanInvestorFundByID(context.Context, int32) error

	// From loan_investor_funds_fund_id_idx
	LoanInvestorFundByFundID(context.Context, string) (*LoanInvestorFund, error)

	UpdateLoanInvestorFundByFundID(context.Context, *LoanInvestorFund) error
	DeleteLoanInvestorFundByFundID(context.Context, string) error
}
