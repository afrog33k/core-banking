// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package roles

import (
	"context"

	pb "bnk.to/core/api/v1/roles"
	"bnk.to/core/tools/db/mux"
)

func (s *UsersServer) GetUserRole(ctx context.Context, req *pb.GetUserRoleRequest) (*pb.UserRole, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.roles.users.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.RoleUserRoleByRoleID(ctx, req.RoleID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
