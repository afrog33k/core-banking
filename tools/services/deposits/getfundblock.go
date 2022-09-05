// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package deposits

import (
	"context"

	pb "bnk.to/core/api/v1/deposits"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetFundBlock(ctx context.Context, req *pb.GetFundBlockRequest) (*pb.FundBlock, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.deposits.blocks.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.DepositFundBlockByBlockID(ctx, req.BlockID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
