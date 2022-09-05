// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package loans

import (
	"context"

	pb "bnk.to/core/api/v1/loans"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) CreateFunding(ctx context.Context, req *pb.CreateFundingRequest) (*pb.InvestorFund, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.fundings.create"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := db.NewLoanInvestorFund(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.InsertLoanInvestorFund(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}
