package main

import (
	"context"
	"fmt"
	"github.com/threpio/shed/data"
)

type HParentOrgCreateBody struct {
	Name string
}

type HParentOrgCreateResponse struct {
	ID   int64
	Name string
}

type HParentOrgListResponse struct {
	ParentOrganisation []data.ParentOrganisation
}

type IParentOrg struct {
	query *data.Queries
}

func NewIParentOrg(query *data.Queries) *IParentOrg {
	return &IParentOrg{
		query: query,
	}
}

func (p *IParentOrg) Create(ctx context.Context, B HParentOrgCreateBody) (*HParentOrgCreateResponse, error) {
	//TODO: Decide whether the names can be duplicated
	insertedParentOrg, err := p.query.InsertParentOrganisation(ctx, B.Name)
	if err != nil {
		return nil, err
	}
	fmt.Println(insertedParentOrg)

	insertedParentOrgResponse := &HParentOrgCreateResponse{
		ID:   insertedParentOrg.ID,
		Name: insertedParentOrg.Name,
	}

	return insertedParentOrgResponse, nil
}

func (p *IParentOrg) List(ctx context.Context) (*HParentOrgListResponse, error) {
	print("List")
	list, err := p.query.ListParentOrganisations(ctx)
	if err != nil {
		return nil, err
	}

	listResponse := &HParentOrgListResponse{
		ParentOrganisation: list,
	}

	return listResponse, nil
}
