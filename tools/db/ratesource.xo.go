// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"bnk.to/core/api/v1/rates"
)

// RateSource represents a row from 'rate_sources'.
type RateSource struct {
	ID       int32           `json:"id"`        // id
	SourceID string          `json:"source_id"` // source_id
	Name     string          `json:"name"`      // name
	Type     RatesSourceType `json:"type"`      // type
	Notes    string          `json:"notes"`     // notes
	// xo fields
	Exists, Deleted bool
}

func NewRateSource(pb *rates.Source) (RateSource, error) {
	if pb == nil {
		return RateSource{}, ErrNilType{"RateSource"}
	}
	rs := RateSource{
		SourceID: pb.SourceID,
		Name:     pb.Name,
		Type:     NewRatesSourceType(pb.Type),
		Notes:    pb.Notes,
	}
	return rs, nil
}

func (rs RateSource) PB() (*rates.Source, error) {
	pb := &rates.Source{
		SourceID: rs.SourceID,
		Name:     rs.Name,
		Type:     rs.Type.PB(),
		Notes:    rs.Notes,
	}
	return pb, nil
}

type RateSourceRepository interface {
	InsertRateSource(context.Context, *RateSource) error
	ListRateSources(context.Context, string, int32, string, *ListPosition) (ListStat, []*RateSource, *ListPosition, error)

	// From rate_sources_pkey
	RateSourceByID(context.Context, int32) (*RateSource, error)

	UpdateRateSourceByID(context.Context, *RateSource) error
	DeleteRateSourceByID(context.Context, int32) error

	// From rate_sources_source_id_idx
	RateSourceBySourceID(context.Context, string) (*RateSource, error)

	UpdateRateSourceBySourceID(context.Context, *RateSource) error
	DeleteRateSourceBySourceID(context.Context, string) error
}
