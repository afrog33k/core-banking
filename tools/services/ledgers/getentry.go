// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package ledgers

import (
	"context"

	pb "bnk.to/core/api/v1/ledgers"
	"bnk.to/core/tools/db/mux"
)

func (s *EntriesServer) GetEntry(ctx context.Context, req *pb.GetEntryRequest) (*pb.Entry, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.ledgers.entries.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.LedgerEntryByEntryID(ctx, req.EntryID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
