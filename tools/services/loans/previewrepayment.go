// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package loans

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/loans"
)

func (s *Server) PreviewRepayment(ctx context.Context, req *pb.PreviewRepaymentRequest) (*pb.PreviewRepaymentResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.previewrepayment"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
