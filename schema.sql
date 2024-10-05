CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE competitions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    instance VARCHAR(255) NOT NULL
);

CREATE TABLE competition_days (
    id SERIAL PRIMARY KEY,
    competition_id INT NOT NULL,
    day VARCHAR(255) NOT NULL,
    FOREIGN KEY (competition_id) REFERENCES competitions(id)
);

CREATE TABLE parent_organisations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TYPE PLAY_CATEGORIES AS ENUM ('Open', 'Mixed', 'WomenPlus', 'Other');

CREATE TABLE refs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_organisation INT NOT NULL,
    play_category PLAY_CATEGORIES NOT NULL,
    competition_days INT NOT NULL,
    FOREIGN KEY (parent_organisation) REFERENCES parent_organisations(id),
    FOREIGN KEY (competition_days) REFERENCES competition_days(id)
);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    competition_id INT NOT NULL,
    parent_organisation_id INT NOT NULL,
    play_category PLAY_CATEGORIES NOT NULL,
    FOREIGN KEY (parent_organisation_id) REFERENCES parent_organisations(id),
    FOREIGN KEY (competition_id) REFERENCES competitions(id)
);

CREATE TABLE team_participation (
    team_id INT NOT NULL,
    competition_day_id INT NOT NULL,
    PRIMARY KEY (team_id, competition_day_id),
    FOREIGN KEY (team_id) REFERENCES teams(id),
    FOREIGN KEY (competition_day_id) REFERENCES competition_days(id)
);

CREATE TABLE Rounds (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    number INT NOT NULL,
    competition_day_id INT NOT NULL,
    FOREIGN KEY (competition_day_id) REFERENCES competition_days(id)
);

CREATE TABLE Matches (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    round_uuid UUID NOT NULL,
    refs_id INT NOT NULL,
    team1_id INT NOT NULL,
    team2_id INT NOT NULL,
    court INT NOT NULL,
    FOREIGN KEY (refs_id) REFERENCES refs(id),
    FOREIGN KEY (round_uuid) REFERENCES Rounds(uuid),
    FOREIGN KEY (team1_id) REFERENCES teams(id),
    FOREIGN KEY (team2_id) REFERENCES teams(id)
);