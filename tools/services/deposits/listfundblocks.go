// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package deposits

import (
	"context"

	pb "bnk.to/core/api/v1/deposits"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
	"bnk.to/core/tools/services"
)

func (s *Server) ListFundBlocks(ctx context.Context, req *pb.ListFundBlocksRequest) (*pb.ListFundBlocksResponse, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.deposits.blocks.list"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	var after *db.ListPosition
	if req.PageToken != "" {
		if req.Filter != "" {
			return nil, services.ErrFilterOnPagination
		}
		if req.PageSize != 0 {
			return nil, services.ErrSizeOnPagination
		}
		if req.OrderBy != "" {
			return nil, services.ErrOrderOnPagination
		}
		var err error
		req.Filter, req.PageSize, req.OrderBy, after, err = services.DecodePageToken(req.PageToken)
		if err != nil {
			return nil, err
		}
	}
	if req.PageSize == 0 {
		req.PageSize = services.DefaultPageSize
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		return nil, services.ErrInvalidPageSize
	}
	stat, v, newAfter, err := storage.ListDepositFundBlocks(ctx, req.Filter, req.PageSize, req.OrderBy, after)
	if err != nil {
		return nil, err
	}
	results := make([]*pb.FundBlock, 0, len(v))
	for _, w := range v {
		pb, err := w.PB()
		if err != nil {
			return nil, err
		}
		results = append(results, pb)
	}
	return &pb.ListFundBlocksResponse{
		Total:     stat.Total,
		Remaining: stat.Remaining,
		Blocks:    results,
		NextPageToken: services.EncodePageToken(
			req.Filter, req.PageSize, req.OrderBy, *newAfter,
		),
	}, nil
}
