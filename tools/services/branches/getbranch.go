// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package branches

import (
	"context"

	pb "bnk.to/core/api/v1/branches"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetBranch(ctx context.Context, req *pb.GetBranchRequest) (*pb.Branch, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.branches.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.BranchByBranchID(ctx, req.BranchID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
