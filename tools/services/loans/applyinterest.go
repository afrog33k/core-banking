// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package loans

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/loans"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ApplyInterest(ctx context.Context, req *pb.ApplyInterestRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.applyinterest"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
