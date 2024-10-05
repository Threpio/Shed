package main

import (
	"context"
	"github.com/threpio/shed/data"
	"log"
	_ "reflect"

	"github.com/jackc/pgx/v5"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=theoa dbname=postgres sslmode=verify-full")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := data.New(conn)

	// Create a Competition
	insertedCompetition, err := queries.InsertCompetition(ctx, data.InsertCompetitionParams{
		Name:     "Test Competition",
		Instance: "1",
	})
	if err != nil {
		return err
	}
	log.Println(insertedCompetition)

	// List all Competitions
	competitions, err := queries.ListCompetitions(ctx)
	if err != nil {
		return err
	}
	log.Println(competitions)

	// Create a Competition Day
	insertedCompetitionDay, err := queries.InsertCompetitionDay(ctx, data.InsertCompetitionDayParams{
		CompetitionID: insertedCompetition.ID,
		Day:           "2021-06-01",
	})
	if err != nil {
		return err
	}
	log.Println(insertedCompetitionDay)

	// Create a Parent Organisation
	insertedParentOrganisation, err := queries.InsertParentOrganisation(ctx, "Lugi Test")
	if err != nil {
		return err
	}
	log.Println(insertedParentOrganisation)

	// List all Parent Organisations
	parentOrganisations, err := queries.ListParentOrganisations(ctx)
	if err != nil {
		return err
	}
	log.Println(parentOrganisations)

	// Create a Team
	insertedTeam, err := queries.InsertTeam(ctx, data.InsertTeamParams{
		Name:                 "Lugi Mixed",
		CompetitionID:        insertedCompetition.ID,
		ParentOrganisationID: parentOrganisations[0].ID,
		PlayCategory:         data.PlayCategoriesMixed,
	})
	if err != nil {
		return err
	}
	log.Println(insertedTeam)

	insertedTeam2, err := queries.InsertTeam(ctx, data.InsertTeamParams{
		Name:                 "Lugi Mixed 2",
		CompetitionID:        insertedCompetition.ID,
		ParentOrganisationID: parentOrganisations[0].ID,
		PlayCategory:         data.PlayCategoriesMixed,
	})

	insertedReferee, err := queries.InsertRef(ctx, data.InsertRefParams{
		Name:               "Lugi Mens",
		ParentOrganisation: insertedParentOrganisation.ID,
		PlayCategory:       data.PlayCategoriesOpen,
		CompetitionDays:    insertedCompetitionDay.ID,
	})

	// List all Teams
	teams, err := queries.ListTeams(ctx)
	if err != nil {
		return err
	}
	log.Println(teams)

	// Create a round
	insertedRound, err := queries.InsertRound(ctx, data.InsertRoundParams{
		Number:           1,
		CompetitionDayID: insertedCompetitionDay.ID,
	})
	if err != nil {
		return err
	}
	log.Println(insertedRound)

	// Create a match
	insertedMatch, err := queries.InsertMatch(ctx, data.InsertMatchParams{
		RoundUuid: insertedRound.Uuid,
		RefsID:    insertedReferee.ID,
		Team1ID:   insertedTeam.ID,
		Team2ID:   insertedTeam2.ID,
		Court:     1,
	})
	if err != nil {
		return err
	}
	log.Println(insertedMatch)

	// List all Matches
	matches, err := queries.ListMatches(ctx)
	if err != nil {
		return err
	}
	log.Println(matches)

	// GetMatchByCompetitionDayID
	matchesByCompetitionDayID, err := queries.GetMatchByCompetitionDayID(ctx, insertedCompetitionDay.ID)
	if err != nil {
		return err
	}
	log.Println(matchesByCompetitionDayID)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
