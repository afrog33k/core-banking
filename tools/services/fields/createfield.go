// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package fields

import (
	"context"

	pb "bnk.to/core/api/v1/fields"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) CreateField(ctx context.Context, req *pb.CreateFieldRequest) (*pb.Field, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.fields.fields.create"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := db.NewField(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.InsertField(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}
