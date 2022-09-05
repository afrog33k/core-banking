// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/groups"
)

// Group represents a row from 'groups'.
type Group struct {
	ID               int32       `json:"id"`                 // id
	GroupID          string      `json:"group_id"`           // group_id
	Info             ContactInfo `json:"info"`               // info
	BranchID         string      `json:"branch_id"`          // branch_id
	CentreID         string      `json:"centre_id"`          // centre_id
	OwnerID          string      `json:"owner_id"`           // owner_id
	LoanCycle        int32       `json:"loan_cycle"`         // loan_cycle
	MigrationEventID string      `json:"migration_event_id"` // migration_event_id
	Notes            string      `json:"notes"`              // notes
	CreateTime       time.Time   `json:"create_time"`        // create_time
	UpdateTime       time.Time   `json:"update_time"`        // update_time
	// xo fields
	Exists, Deleted bool
}

func NewGroup(pb *groups.Group) (Group, error) {
	if pb == nil {
		return Group{}, ErrNilType{"Group"}
	}
	g := Group{
		GroupID:          pb.GroupID,
		BranchID:         pb.BranchID,
		CentreID:         pb.CentreID,
		OwnerID:          pb.OwnerID,
		LoanCycle:        pb.LoanCycle,
		MigrationEventID: pb.MigrationEventID,
		Notes:            pb.Notes,
		CreateTime:       pb.CreateTime.AsTime(),
		UpdateTime:       pb.UpdateTime.AsTime(),
	}
	var err error
	g.Info, err = NewContactInfo(pb.Info)
	if err != nil {
		return Group{}, err
	}
	return g, nil
}

func (g Group) PB() (*groups.Group, error) {
	pb := &groups.Group{
		GroupID:          g.GroupID,
		BranchID:         g.BranchID,
		CentreID:         g.CentreID,
		OwnerID:          g.OwnerID,
		LoanCycle:        g.LoanCycle,
		MigrationEventID: g.MigrationEventID,
		Notes:            g.Notes,
		CreateTime:       timestamppb.New(g.CreateTime),
		UpdateTime:       timestamppb.New(g.UpdateTime),
	}
	var err error
	pb.Info, err = g.Info.PB()
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type GroupRepository interface {
	InsertGroup(context.Context, *Group) error
	ListGroups(context.Context, string, int32, string, *ListPosition) (ListStat, []*Group, *ListPosition, error)

	// From groups_pkey
	GroupByID(context.Context, int32) (*Group, error)

	UpdateGroupByID(context.Context, *Group) error
	DeleteGroupByID(context.Context, int32) error

	// From groups_group_id_idx
	GroupByGroupID(context.Context, string) (*Group, error)

	UpdateGroupByGroupID(context.Context, *Group) error
	DeleteGroupByGroupID(context.Context, string) error
}
