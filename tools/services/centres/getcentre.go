// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package centres

import (
	"context"

	pb "bnk.to/core/api/v1/centres"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetCentre(ctx context.Context, req *pb.GetCentreRequest) (*pb.Centre, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.centres.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.CentreByCentreID(ctx, req.CentreID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
