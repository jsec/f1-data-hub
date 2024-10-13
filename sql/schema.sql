-- We need this for sqlc to reference
-- TODO: figure out a better way to do this

CREATE TABLE circuits (
    id integer NOT NULL,
    ref character varying(255) DEFAULT ''::character varying NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    location character varying(255),
    country character varying(255),
    lat double precision,
    lng double precision,
    alt integer,
    url character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE TABLE constructor_results (
    id integer NOT NULL,
    race_id integer DEFAULT 0 NOT NULL,
    constructor_id integer DEFAULT 0 NOT NULL,
    points double precision,
    status character varying(255)
);
CREATE TABLE constructor_standings (
    id integer NOT NULL,
    race_id integer DEFAULT 0 NOT NULL,
    constructor_id integer DEFAULT 0 NOT NULL,
    points double precision DEFAULT '0'::double precision NOT NULL,
    position integer,
    position_text character varying(255),
    wins integer DEFAULT 0 NOT NULL
);

CREATE TABLE constructors (
    id integer NOT NULL,
    ref character varying(255) DEFAULT ''::character varying NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    nationality character varying(255),
    url character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE TABLE driver_standings (
    id integer NOT NULL,
    race_id integer DEFAULT 0 NOT NULL,
    driver_id integer DEFAULT 0 NOT NULL,
    points double precision DEFAULT '0'::double precision NOT NULL,
    position integer,
    position_text character varying(255),
    wins integer DEFAULT 0 NOT NULL
);

CREATE TABLE drivers (
    id integer NOT NULL,
    ref character varying(255) DEFAULT ''::character varying NOT NULL,
    number integer,
    code character varying(3),
    first_name character varying(255) DEFAULT ''::character varying NOT NULL,
    last_name character varying(255) DEFAULT ''::character varying NOT NULL,
    date_of_birth date,
    nationality character varying(255),
    url character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE TABLE lap_times (
    race_id integer NOT NULL,
    driver_id integer NOT NULL,
    lap integer NOT NULL,
    position integer,
    time character varying(255),
    milliseconds integer
);

CREATE TABLE pit_stops (
    race_id integer NOT NULL,
    driver_id integer NOT NULL,
    stop integer NOT NULL,
    lap integer NOT NULL,
    time time without time zone NOT NULL,
    duration character varying(255),
    milliseconds integer
);

CREATE TABLE qualifying (
    id integer NOT NULL,
    race_id integer DEFAULT 0 NOT NULL,
    driver_id integer DEFAULT 0 NOT NULL,
    constructor_id integer DEFAULT 0 NOT NULL,
    number integer DEFAULT 0 NOT NULL,
    position integer,
    q1 character varying(255),
    q2 character varying(255),
    q3 character varying(255)
);

CREATE TABLE races (
    id integer NOT NULL,
    year integer DEFAULT 0 NOT NULL,
    round integer DEFAULT 0 NOT NULL,
    circuit_id integer DEFAULT 0 NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    date date NOT NULL,
    time time without time zone,
    url character varying(255),
    fp1_date date,
    fp1_time time without time zone,
    fp2_date date,
    fp2_time time without time zone,
    fp3_date date,
    fp3_time time without time zone,
    quali_date date,
    quali_time time without time zone,
    sprint_date date,
    sprint_time time without time zone
);

CREATE TABLE results (
    id integer NOT NULL,
    race_id integer DEFAULT 0 NOT NULL,
    driver_id integer DEFAULT 0 NOT NULL,
    constructor_id integer DEFAULT 0 NOT NULL,
    number integer,
    grid integer DEFAULT 0 NOT NULL,
    position integer,
    position_text character varying(255) DEFAULT ''::character varying NOT NULL,
    position_order integer DEFAULT 0 NOT NULL,
    points double precision DEFAULT '0'::double precision NOT NULL,
    laps integer DEFAULT 0 NOT NULL,
    time character varying(255),
    milliseconds integer,
    fastest_lap integer,
    rank integer DEFAULT 0,
    fastest_lap_time character varying(255),
    fastest_lap_speed character varying(255),
    status_id integer DEFAULT 0 NOT NULL
);

CREATE TABLE seasons (
    year integer DEFAULT 0 NOT NULL,
    url character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE TABLE sprint_results (
    id integer NOT NULL,
    race_id integer DEFAULT 0 NOT NULL,
    driver_id integer DEFAULT 0 NOT NULL,
    constructor_id integer DEFAULT 0 NOT NULL,
    number integer DEFAULT 0 NOT NULL,
    grid integer DEFAULT 0 NOT NULL,
    position integer,
    position_text character varying(255) DEFAULT ''::character varying NOT NULL,
    position_order integer DEFAULT 0 NOT NULL,
    points double precision DEFAULT '0'::double precision NOT NULL,
    laps integer DEFAULT 0 NOT NULL,
    time character varying(255),
    milliseconds integer,
    fastest_lap integer,
    fastest_lap_time character varying(255),
    status_id integer DEFAULT 0 NOT NULL
);

CREATE TABLE status (
    id integer NOT NULL,
    status character varying(255) DEFAULT ''::character varying NOT NULL
);
