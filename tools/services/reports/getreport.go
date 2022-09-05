// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package reports

import (
	"context"

	pb "bnk.to/core/api/v1/reports"
	"bnk.to/core/tools/db/mux"
)

func (s *Server) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.Report, error) {
	if err := s.Auth.CheckPerm(ctx, req, "v1.reports.get"); err != nil {
		return nil, err
	}

	storage := mux.Storage(ctx)
	v, err := storage.ReportByReportID(ctx, req.ReportID)
	if err != nil {
		return nil, err
	}
	return v.PB()
}
