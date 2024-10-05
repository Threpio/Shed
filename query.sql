-- name: InsertCompetition :one
INSERT INTO competitions (
    name, instance
) VALUES (?, ?)
RETURNING *;

-- name: ListCompetitions :many
SELECT * FROM competitions limit 100;

-- name: GetCompetitionByName :one
SELECT * FROM competitions
WHERE name = ? LIMIT 1;

-- name: GetCompetitionByID :one
SELECT * FROM competitions
WHERE id = ? LIMIT 1;

-- name: UpdateCompetitions :exec
UPDATE competitions SET
    name = ?,
    instance = ?
WHERE id = ?;



-- name: InsertCompetitionDay :one
INSERT INTO competition_days (
    competition_id, day
) VALUES (?, ?)
RETURNING *;

-- name: ListCompetitionDays :many
SELECT * FROM competition_days limit 100;

-- name: ListCompetitionDaysByCompetitionID :many
SELECT * FROM competition_days
WHERE competition_id = ? LIMIT 100;

-- name: GetCompetitionDayByID :one
SELECT * FROM competition_days
WHERE id = ? LIMIT 1;

-- name: UpdateCompetitionDay :exec
UPDATE competition_days SET
    competition_id = ?,
    day = ?
WHERE id = ?;



-- name: InsertParentOrganisation :one
INSERT INTO parent_organisations (
    name
) VALUES (?)
RETURNING *;

-- name: ListParentOrganisations :many
SELECT * FROM parent_organisations limit 100;

-- name: GetParentOrganisationByID :one
SELECT * FROM parent_organisations
WHERE id = ? LIMIT 1;

-- name: GetParentOrganisationByName :one
SELECT * FROM parent_organisations
WHERE name = ? LIMIT 1;

-- name: SearchParentOrganisationByName :many
SELECT * FROM parent_organisations
WHERE name LIKE '%' || ? || '%' limit 100;

-- name: UpdateParentOrganisation :exec
UPDATE parent_organisations SET
    name = ?
WHERE id = ?;



-- name: InsertRef :one
INSERT INTO refs (
    name, parent_organisation, play_category, competition_days
) VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListRefs :many
SELECT * FROM refs limit 100;

-- name: ListRefsByCompetitonDayID :many
SELECT * FROM refs
WHERE competition_days = ? limit 100;

-- name: GetRefByID :one
SELECT * FROM refs
WHERE id = ? LIMIT 1;

-- name: UpdateRef :exec
UPDATE refs SET
    name = ?,
    parent_organisation = ?,
    play_category = ?,
    competition_days = ?
WHERE id = ?;



-- name: InsertTeam :one
INSERT INTO teams (
    name, competition_id, parent_organisation_id, play_category
) VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListTeams :many
SELECT * FROM teams limit 100;

-- name: ListTeamsByCompetitionID :many
SELECT * FROM teams
WHERE competition_id = ? limit 100;

-- name: ListTeamsByParentOrganisationID :many
SELECT * FROM teams
WHERE parent_organisation_id = ? limit 100;

-- name: GetTeamByID :one
SELECT * FROM teams
WHERE id = ? LIMIT 1;

-- name: UpdateTeam :exec
UPDATE teams SET
    name = ?,
    competition_id = ?,
    parent_organisation_id = ?,
    play_category = ?
WHERE id = ?;



-- name: InsertTeamParticipation :one
INSERT INTO team_participation (
    team_id, competition_day_id
) VALUES (?, ?)
RETURNING *;

-- name: ListTeamParticipation :many
SELECT * FROM team_participation limit 100;

-- name: ListTeamParticipationByTeamID :many
SELECT * FROM team_participation
WHERE team_id = ? limit 100;

-- name: ListTeamParticipationByCompetitionDayID :many
SELECT * FROM team_participation
WHERE competition_day_id = ? limit 100;

-- name: UpdateTeamParticipation :exec
UPDATE team_participation SET
    team_id = ?,
    competition_day_id = ?
WHERE team_id = ? AND competition_day_id = ?;

-- name: DeleteTeamParticipation :exec
DELETE FROM team_participation
WHERE team_id = ? AND competition_day_id = ?;



-- name: InsertRound :one
INSERT INTO rounds (
    number, competition_day_id
) VALUES (?, ?)
RETURNING *;

-- name: ListRounds :many
SELECT * FROM rounds limit 100;

-- name: ListRoundsByCompetitionDayID :many
SELECT * FROM rounds
WHERE competition_day_id = ? limit 100;

-- name: GetRoundByID :one
SELECT * FROM rounds
WHERE uuid = ? LIMIT 1;

-- name: UpdateRound :exec
UPDATE rounds SET
    number = ?,
    competition_day_id = ?
WHERE uuid = ?;

-- name: DeleteRound :exec
DELETE FROM rounds
WHERE uuid = ?;



-- name: InsertMatch :one
INSERT INTO matches (
    round_uuid, refs_id, team1_id, team2_id, court)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: ListMatches :many
SELECT * FROM matches limit 100;

-- name: ListMatchesByRoundID :many
SELECT * FROM matches
WHERE round_uuid = ? limit 100;

-- name: GetMatchByID :one
SELECT * FROM matches
WHERE uuid = ? LIMIT 1;






-- name: GetMatchByCompetitionDayID :many
SELECT
    M.uuid AS match_id,
    M.team1_id,
    T1.name AS team1_name,
    M.team2_id,
    T2.name AS team2_name,
    M.court,
    M.refs_id,
    R.name AS ref_name,
    R.play_category,
    CD.day AS competition_day,
    RD.number AS round_number
FROM matches M
         JOIN rounds RD ON M.round_uuid = RD.uuid
         JOIN competition_days CD ON RD.competition_day_id = CD.id
         JOIN teams T1 ON M.team1_id = T1.id
         JOIN teams T2 ON M.team2_id = T2.id
         JOIN refs R ON M.refs_id = R.id
WHERE CD.id = ?;

-- name: GetTeamParticipationByCompetitionDayID :many
SELECT
    T.id AS team_id,
    T.name AS team_name,
    PO.name AS parent_organisation_name,
    T.play_category,
    CD.day AS competition_day
FROM team_participation TP
         JOIN teams T ON TP.team_id = T.id
         JOIN competition_days CD ON TP.competition_day_id = CD.id
         JOIN parent_organisations PO ON T.parent_organisation_id = PO.id
WHERE CD.id = ?;











