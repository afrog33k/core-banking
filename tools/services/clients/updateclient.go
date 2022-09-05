// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package clients

import (
	"context"
	"time"

	pb "bnk.to/core/api/v1/clients"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) UpdateClient(ctx context.Context, req *pb.UpdateClientRequest) (*pb.Client, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.clients.update"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	// Override UpdateTime to the current time.
	req.Body.UpdateTime = timestamppb.New(time.Now())
	v, err := db.NewClient(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.UpdateClientByClientID(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}
