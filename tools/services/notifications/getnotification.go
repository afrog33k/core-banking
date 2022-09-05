// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package notifications

import (
	"context"

	pb "bnk.to/core/api/v1/notifications"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetNotification(ctx context.Context, req *pb.GetNotificationRequest) (*pb.Notification, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.notifications.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.NotificationByNotificationID(ctx, req.NotificationID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
