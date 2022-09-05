// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package products

import (
	"context"

	pb "bnk.to/core/api/v1/products"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
)

func (s *RiskLevelsServer) UpdateRiskLevel(ctx context.Context, req *pb.UpdateRiskLevelRequest) (*pb.RiskLevel, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.risklevels.update"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := db.NewProductRiskLevel(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.UpdateProductRiskLevelByLevelID(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}
