-- +goose Up

CREATE TABLE circuits (
    id integer PRIMARY KEY,
    ref text NOT NULL,
    name text NOT NULL,
    location text DEFAULT NULL,
    country text DEFAULT NULL,
    lat numeric DEFAULT NULL,
    lng numeric DEFAULT NULL,
    alt integer DEFAULT NULL,
    url text NOT NULL UNIQUE
);

CREATE TABLE status (
    id integer NOT NULL PRIMARY KEY,
    status text NOT NULL DEFAULT ''
);

CREATE TABLE seasons (
    year integer NOT NULL,
    url text NOT NULL UNIQUE
);

CREATE TABLE constructors (
    id integer PRIMARY KEY,
    ref text NOT NULL,
    name text NOT NULL UNIQUE,
    nationality text DEFAULT NULL,
    url text NOT NULL
);

CREATE TABLE drivers (
    id integer PRIMARY KEY,
    ref text NOT NULL DEFAULT '',
    number integer DEFAULT NULL,
    code text DEFAULT NULL,
    first_name text NOT NULL DEFAULT '',
    last_name text NOT NULL DEFAULT '',
    date_of_birth date DEFAULT NULL,
    nationality text DEFAULT NULL,
    url text NOT NULL UNIQUE
);

CREATE TABLE races (
    id integer NOT NULL PRIMARY KEY,
    year integer NOT NULL,
    round integer NOT NULL DEFAULT 0,
    circuit_id integer NOT NULL REFERENCES circuits (id),
    name text NOT NULL DEFAULT '',
    date date NOT NULL,
    time time DEFAULT NULL,
    url text DEFAULT NULL UNIQUE,
    fp1_date date DEFAULT NULL,
    fp1_time time DEFAULT NULL,
    fp2_date date DEFAULT NULL,
    fp2_time time DEFAULT NULL,
    fp3_date date DEFAULT NULL,
    fp3_time time DEFAULT NULL,
    quali_date date DEFAULT NULL,
    quali_time time DEFAULT NULL,
    sprint_date date DEFAULT NULL,
    sprint_time time DEFAULT NULL
);

CREATE TABLE constructor_results (
    id integer NOT NULL PRIMARY KEY,
    race_id integer NOT NULL REFERENCES races (id),
    constructor_id integer NOT NULL REFERENCES constructors (id),
    points numeric,
    status text
);

CREATE TABLE constructor_standings (
    id integer PRIMARY KEY,
    race_id integer NOT NULL REFERENCES races (id),
    constructor_id integer NOT NULL REFERENCES constructors (id),
    points numeric NOT NULL,
    position integer DEFAULT NULL,
    pos_text text DEFAULT NULL,
    wins integer NOT NULL DEFAULT 0
);

CREATE TABLE driver_standings (
    id integer PRIMARY KEY,
    race_id integer NOT NULL REFERENCES races (id),
    driver_id integer NOT NULL REFERENCES drivers (id),
    points numeric NOT NULL DEFAULT 0,
    position integer DEFAULT NULL,
    pos_text text DEFAULT NULL,
    wins integer DEFAULT 0
);

CREATE TABLE lap_times (
    race_id integer NOT NULL REFERENCES races (id),
    driver_id integer NOT NULL REFERENCES drivers (id),
    lap integer NOT NULL,
    position integer DEFAULT NULL,
    time text DEFAULT NULL,
    milliseconds integer DEFAULT NULL,
    PRIMARY KEY (race_id, driver_id, lap)
);

CREATE TABLE pit_stops (
    race_id integer NOT NULL REFERENCES races (id),
    driver_id integer NOT NULL REFERENCES drivers (id),
    stop integer NOT NULL,
    lap integer NOT NULL,
    time time NOT NULL,
    duration text DEFAULT NULL,
    milliseconds integer DEFAULT NULL,
    PRIMARY KEY (race_id, driver_id, stop)
);

CREATE TABLE qualifying (
    id integer NOT NULL PRIMARY KEY,
    race_id integer NOT NULL REFERENCES races (id),
    driver_id integer NOT NULL REFERENCES drivers (id),
    constructor_id integer NOT NULL REFERENCES constructors (id),
    number integer NOT NULL DEFAULT 0,
    position integer DEFAULT NULL,
    q1 text DEFAULT NULL,
    q2 text DEFAULT NULL,
    q3 text DEFAULT NULL
);

CREATE TABLE results (
    id integer NOT NULL PRIMARY KEY,
    race_id integer NOT NULL REFERENCES races (id),
    driver_id integer NOT NULL REFERENCES drivers (id),
    constructor_id integer NOT NULL REFERENCES constructors (id),
    number integer DEFAULT NULL,
    grid integer NOT NULL DEFAULT 0,
    position integer DEFAULT NULL,
    pos_text text NOT NULL DEFAULT '',
    pos_order integer NOT NULL DEFAULT 0,
    points numeric NOT NULL DEFAULT 0,
    laps integer NOT NULL DEFAULT 0,
    time text DEFAULT NULL,
    milliseconds integer DEFAULT NULL,
    fastest_lap integer DEFAULT NULL,
    rank integer DEFAULT 0,
    fastest_lap_time text DEFAULT NULL,
    fastest_lap_speed text DEFAULT NULL,
    status_id integer REFERENCES status (id)
);

CREATE TABLE sprint_results (
    id integer NOT NULL PRIMARY KEY,
    race_id integer NOT NULL REFERENCES races (id),
    driver_id integer NOT NULL REFERENCES drivers (id),
    constructor_id integer NOT NULL REFERENCES constructors (id),
    number integer NOT NULL DEFAULT 0,
    grid integer NOT NULL DEFAULT 0,
    position integer DEFAULT NULL,
    pos_text text NOT NULL DEFAULT '',
    pos_order integer NOT NULL DEFAULT 0,
    points numeric NOT NULL DEFAULT 0,
    laps integer NOT NULL DEFAULT 0,
    time text DEFAULT NULL,
    milliseconds integer DEFAULT NULL,
    fastest_lap integer DEFAULT NULL,
    fastest_lap_time text DEFAULT NULL,
    status_id integer REFERENCES status (id)
);

-- +goose Down
DROP DATABASE 
