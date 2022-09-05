// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package clients

import (
	"context"
	"fmt"

	pb "bnk.to/core/api/v1/clients"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) RemoveClientDocuments(ctx context.Context, req *pb.RemoveClientDocumentsRequest) (*emptypb.Empty, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.clients.documents"); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("unimplemented")
}
