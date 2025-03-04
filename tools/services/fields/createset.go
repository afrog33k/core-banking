// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package fields

import (
	"context"
	"time"

	pb "bnk.to/core/api/v1/fields"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateSet(ctx context.Context, req *pb.CreateSetRequest) (*pb.Set, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.fields.create"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	req.Body.CreateTime = timestamppb.New(time.Now())
	req.Body.UpdateTime = timestamppb.New(time.Now())
	v, err := db.NewFieldSet(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.InsertFieldSet(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}
