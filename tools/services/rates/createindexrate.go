// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package rates

import (
	"context"

	pb "bnk.to/core/api/v1/rates"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) CreateIndexRate(ctx context.Context, req *pb.CreateIndexRateRequest) (*pb.IndexRate, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.sources.indexrates.create"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := db.NewRateIndexRate(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.InsertRateIndexRate(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}
