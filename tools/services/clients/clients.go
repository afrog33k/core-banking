// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package clients

import (
	"bnk.to/core/tools/services"

	pb "bnk.to/core/api/v1/clients"
)

// Server is the implementation of ClientsService.
type Server struct {
	pb.UnsafeClientsServiceServer
	services.Common
}

// NewServer creates a new Server.
func NewServer(common services.Common) *Server {
	return &Server{
		Common: common,
	}
}
