package main

import "github.com/threpio/shed/data"

type HTeamCreateBody struct {
	Name string `json:"name"`
}

type ITeam struct {
	query data.Queries
}
