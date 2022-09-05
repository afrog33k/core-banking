package transactions

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/transactions"
)

func (s *LoanServer) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.LoanTransaction, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.loans.payment"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
