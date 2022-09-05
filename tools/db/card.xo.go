// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"bnk.to/core/api/v1/cards"
)

// Card represents a row from 'cards'.
type Card struct {
	ID        int32     `json:"id"`         // id
	CardID    string    `json:"card_id"`    // card_id
	Type      CardsType `json:"type"`       // type
	AccountID string    `json:"account_id"` // account_id
	// xo fields
	Exists, Deleted bool
}

func NewCard(pb *cards.Card) (Card, error) {
	if pb == nil {
		return Card{}, ErrNilType{"Card"}
	}
	c := Card{
		CardID:    pb.CardID,
		Type:      NewCardsType(pb.Type),
		AccountID: pb.AccountID,
	}
	return c, nil
}

func (c Card) PB() (*cards.Card, error) {
	pb := &cards.Card{
		CardID:    c.CardID,
		Type:      c.Type.PB(),
		AccountID: c.AccountID,
	}
	return pb, nil
}

type CardRepository interface {
	InsertCard(context.Context, *Card) error
	ListCards(context.Context, string, int32, string, *ListPosition) (ListStat, []*Card, *ListPosition, error)

	// From cards_pkey
	CardByID(context.Context, int32) (*Card, error)

	UpdateCardByID(context.Context, *Card) error
	DeleteCardByID(context.Context, int32) error

	// From cards_card_id_idx
	CardByCardID(context.Context, string) (*Card, error)

	UpdateCardByCardID(context.Context, *Card) error
	DeleteCardByCardID(context.Context, string) error

	// From cards_type_idx
	CardByType(context.Context, CardsType) ([]*Card, error)
}
