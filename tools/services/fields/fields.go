// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package fields

import (
	"bnk.to/core/tools/services"

	pb "bnk.to/core/api/v1/fields"
)

// Server is the implementation of FieldsService.
type Server struct {
	pb.UnsafeFieldsServiceServer
	services.Common
}

// NewServer creates a new Server.
func NewServer(common services.Common) *Server {
	return &Server{
		Common: common,
	}
}
