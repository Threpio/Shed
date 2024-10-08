package main

import (
	"context"
	"fmt"
	"github.com/threpio/shed/data"
)

type HRefCreateBody struct {
	Name               string
	ParentOrganisation int64
	PlayCategory       string
	CompetitionDays    int64
}

type HRefCreateResponse struct {
	ID                 int64
	Name               string
	ParentOrganisation int64
	PlayCategory       string
	CompetitionDays    int64
}

type HRefListResponse struct {
	Refs []data.Ref
}

type IRefs struct {
	query *data.Queries
}

func NewIRefs(query *data.Queries) *IRefs {
	return &IRefs{
		query: query,
	}
}

func (r *IRefs) Create(ctx context.Context, B HRefCreateBody) (*HRefCreateResponse, error) {

	insertedRef, err := r.query.InsertRef(ctx, data.InsertRefParams{
		Name:               B.Name,
		ParentOrganisation: B.ParentOrganisation,
		PlayCategory:       B.PlayCategory,
		CompetitionDays:    B.CompetitionDays,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(insertedRef)

	insertedRefResponse := &HRefCreateResponse{
		ID:                 insertedRef.ID,
		Name:               insertedRef.Name,
		ParentOrganisation: insertedRef.ParentOrganisation,
		PlayCategory:       insertedRef.PlayCategory,
		CompetitionDays:    insertedRef.CompetitionDays,
	}

	return insertedRefResponse, nil
}

func (r *IRefs) List(ctx context.Context) (*HRefListResponse, error) {
	print("List")
	list, err := r.query.ListRefs(ctx)
	if err != nil {
		return nil, err
	}

	listResponse := &HRefListResponse{
		Refs: list,
	}

	return listResponse, nil
}
