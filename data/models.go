// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package data

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Competition struct {
	ID       int64
	Name     string
	Instance string
}

type CompetitionDay struct {
	ID            int64
	CompetitionID int64
	Day           string
}

type Match struct {
	Uuid      pgtype.UUID
	RoundUuid pgtype.UUID
	RefsID    int64
	Team1ID   int64
	Team2ID   int64
	Court     int32
}

type ParentOrganisation struct {
	ID   int64
	Name string
}

type Ref struct {
	ID                 int64
	Name               string
	ParentOrganisation int64
	PlayCategory       string
	CompetitionDays    int64
}

type Round struct {
	Uuid             pgtype.UUID
	Number           int32
	CompetitionDayID int64
}

type Team struct {
	ID                   int64
	Name                 string
	CompetitionID        int64
	ParentOrganisationID int64
	PlayCategory         interface{}
}

type TeamParticipation struct {
	TeamID           int64
	CompetitionDayID int64
}
