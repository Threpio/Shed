CREATE TABLE competitions (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    instance TEXT NOT NULL
);

CREATE TABLE competition_days (
    id BIGSERIAL PRIMARY KEY,
    competition_id BIGINT NOT NULL,
    day TEXT NOT NULL,
    FOREIGN KEY (competition_id) REFERENCES competitions(id)
);

CREATE TABLE parent_organisations (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE refs (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    parent_organisation BIGINT NOT NULL,
    play_category TEXT CHECK( play_category IN ('Mixed','Open','WomensPlus')) NOT NULL,
    competition_days BIGINT NOT NULL,
    FOREIGN KEY (parent_organisation) REFERENCES parent_organisations(id),
    FOREIGN KEY (competition_days) REFERENCES competition_days(id)
);

CREATE TABLE teams (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    competition_id BIGINT NOT NULL,
    parent_organisation_id BIGINT NOT NULL,
    play_category TEXT CHECK( play_category IN ('Mixed','Open','WomensPlus')) NOT NULL,
    FOREIGN KEY (parent_organisation_id) REFERENCES parent_organisations(id),
    FOREIGN KEY (competition_id) REFERENCES competitions(id)
);

CREATE TABLE team_participation (
    team_id BIGINT NOT NULL,
    competition_day_id BIGINT NOT NULL,
    PRIMARY KEY (team_id, competition_day_id),
    FOREIGN KEY (team_id) REFERENCES teams(id),
    FOREIGN KEY (competition_day_id) REFERENCES competition_days(id)
);

CREATE TABLE rounds (
    uuid UUID PRIMARY KEY,
    number INT NOT NULL,
    competition_day_id BIGINT NOT NULL,
    FOREIGN KEY (competition_day_id) REFERENCES competition_days(id)
);

CREATE TABLE matches (
    uuid UUID PRIMARY KEY,
    round_uuid UUID NOT NULL,
    refs_id BIGINT NOT NULL,
    team1_id BIGINT NOT NULL,
    team2_id BIGINT NOT NULL,
    court INT NOT NULL,
    FOREIGN KEY (refs_id) REFERENCES refs(id),
    FOREIGN KEY (round_uuid) REFERENCES Rounds(uuid),
    FOREIGN KEY (team1_id) REFERENCES teams(id),
    FOREIGN KEY (team2_id) REFERENCES teams(id)
);