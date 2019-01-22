CREATE TABLE users (
	id                  serial      PRIMARY KEY,
	created_at          timestamptz NOT NULL,
	deleted_at          timestamptz,
    email               text        NOT NULL UNIQUE,
	profile_image_url   text        NOT NULL,
	pw_hash             bytea,
	pw_iterations       int,
	pw_salt             bytea,
	role                text        NOT NULL,
	updated_at          timestamptz,
	volleynet_user      text,
	volleynet_user_id   int
);

CREATE TABLE tournaments (
	id                  int         PRIMARY KEY,
	created_at          timestamptz NOT NULL,
	updated_at          timestamptz NOT NULL,
    gender              text        NOT NULL,
    signedup_teams      int         NOT NULL,
	start_date          timestamptz NOT NULL,
	end_date            timestamptz NOT NULL,
	name                text        NOT NULL,
    league              text        NOT NULL,
    league_slug         text        NOT NULL,
    sub_league          text        NOT NULL,
    sub_league_slug     text        NOT NULL,
	link                text        NOT NULL,
	entry_link          text        NOT NULL,
	status              text        NOT NULL,
	registration_open   boolean     NOT NULL,
	location            text        NOT NULL,
	live_scoring_link   text        NOT NULL,
	html_notes          text        NOT NULL,
	mode                text        NOT NULL,
    max_points          smallint    NOT NULL,
	min_teams           smallint    NOT NULL,
	max_teams           smallint    NOT NULL,
	end_registration    timestamptz,
	organiser           text        NOT NULL,
	phone               text        NOT NULL,
	email               text        NOT NULL,
	website             text        NOT NULL,
	current_points      text        NOT NULL,
	season              smallint    NOT NULL,
	loc_lat             float8      NOT NULL,
	loc_lon             float8      NOT NULL
);

CREATE INDEX tournaments_name       	ON tournaments USING btree  (name);
CREATE INDEX tournaments_start_date 	ON tournaments USING btree  (start_date);
CREATE INDEX tournaments_end_date   	ON tournaments USING btree  (end_date);
CREATE INDEX tournaments_gender			ON tournaments USING hash   (gender);
CREATE INDEX tournaments_league_slug    ON tournaments USING hash   (league_slug);
CREATE INDEX tournaments_season     	ON tournaments USING hash   (season);

CREATE TABLE players (
	id              int         PRIMARY KEY,
	created_at      timestamptz NOT NULL,
	updated_at      timestamptz NOT NULL,
	first_name      text        NOT NULL,
	last_name       text        NOT NULL,
    total_points    smallint    NOT NULL,
	ladder_rank     smallint    NOT NULL,
	country_union   text        NOT NULL,
	club            text        NOT NULL,
	birthday        date,
	license         text        NOT NULL,
	gender          varchar(1)  NOT NULL
);

CREATE INDEX players_first_name 	ON players USING btree  (first_name);
CREATE INDEX players_last_name  	ON players USING btree  (last_name);
CREATE INDEX players_ladder_rank    ON players USING btree  (ladder_rank);
CREATE INDEX players_gender     	ON players USING hash   (gender);

CREATE TABLE tournament_teams (
	tournament_id   int         NOT NULL REFERENCES tournaments(id),
	player_1_id     int         NOT NULL REFERENCES players(id),
	player_2_id     int         NOT NULL REFERENCES players(id),
	result 			smallint    NOT NULL,
	seed            smallint    NOT NULL,
	total_points    smallint    NOT NULL,
	won_points      smallint    NOT NULL,
	prize_money     float(4)    NOT NULL,
	deregistered    boolean     NOT NULL,

	PRIMARY KEY (tournament_id, player_1_id, player_2_id)
);

CREATE INDEX tournament_teams_team ON tournament_teams (player_1_id, player_2_id);
