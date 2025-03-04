// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package loans

import (
	"context"

	pb "bnk.to/core/api/v1/loans"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"

	cardspb "bnk.to/core/api/v1/cards"
)

func (s *Server) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*cardspb.Card, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.accounts.loans.cards.create"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := db.NewCard(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.InsertCard(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil
}
