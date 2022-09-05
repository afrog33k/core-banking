// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package cards

import (
	"context"

	pb "bnk.to/core/api/v1/cards"
	"bnk.to/core/tools/db/mux"
)

func (s *HoldConfigServer) GetHoldConfig(ctx context.Context, req *pb.GetHoldConfigRequest) (*pb.HoldConfig, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.holds.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.CardHoldConfigByMerchantCode(ctx, req.MerchantCode)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
