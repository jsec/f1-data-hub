-- +goose Up

-- +goose StatementBegin
ALTER TABLE constructorresults RENAME TO constructor_results;
ALTER TABLE constructorstandings RENAME TO constructor_standings;
ALTER TABLE driverstandings RENAME TO driver_standings;
ALTER TABLE laptimes RENAME TO lap_times;
ALTER TABLE pitstops RENAME TO pit_stops;
ALTER TABLE sprintresults RENAME TO sprint_results;

ALTER TABLE circuits RENAME COLUMN circuitid TO id;
ALTER TABLE circuits RENAME COLUMN circuitref TO ref;

ALTER TABLE constructor_results RENAME COLUMN constructorresultsid TO id;
ALTER TABLE constructor_results RENAME COLUMN raceid TO race_id;
ALTER TABLE constructor_results RENAME COLUMN constructorid TO constructor_id;

ALTER TABLE constructors RENAME COLUMN constructorid TO id;
ALTER TABLE constructors RENAME COLUMN constructorref TO ref;

ALTER TABLE constructor_standings RENAME COLUMN constructorstandingsid TO id;
ALTER TABLE constructor_standings RENAME COLUMN raceid TO race_id;
ALTER TABLE constructor_standings RENAME COLUMN constructorid TO constructor_id;
ALTER TABLE constructor_standings RENAME COLUMN positiontext TO position_text;

ALTER TABLE drivers RENAME COLUMN driverid TO id;
ALTER TABLE drivers RENAME COLUMN driverref TO ref;
ALTER TABLE drivers RENAME COLUMN forename TO first_name;
ALTER TABLE drivers RENAME COLUMN surname TO last_name;
ALTER TABLE drivers RENAME COLUMN dob TO date_of_birth;

ALTER TABLE driver_standings RENAME COLUMN driverstandingsid TO id;
ALTER TABLE driver_standings RENAME COLUMN raceid TO race_id;
ALTER TABLE driver_standings RENAME COLUMN driverid TO driver_id;
ALTER TABLE driver_standings RENAME COLUMN positiontext TO position_text;

ALTER TABLE lap_times RENAME COLUMN raceid TO race_id;
ALTER TABLE lap_times RENAME COLUMN driverid TO driver_id;

ALTER TABLE pit_stops RENAME COLUMN raceid TO race_id;
ALTER TABLE pit_stops RENAME COLUMN driverid TO driver_id;

ALTER TABLE qualifying RENAME COLUMN qualifyid TO id;
ALTER TABLE qualifying RENAME COLUMN raceid TO race_id;
ALTER TABLE qualifying RENAME COLUMN driverid TO driver_id;
ALTER TABLE qualifying RENAME COLUMN constructorid TO constructor_id;

ALTER TABLE races RENAME COLUMN raceid TO id;
ALTER TABLE races RENAME COLUMN circuitid TO circuit_id;

ALTER TABLE results RENAME COLUMN resultid TO id;
ALTER TABLE results RENAME COLUMN raceid TO race_id;
ALTER TABLE results RENAME COLUMN driverid TO driver_id;
ALTER TABLE results RENAME COLUMN constructorid TO constructor_id;
ALTER TABLE results RENAME COLUMN positiontext TO position_text;
ALTER TABLE results RENAME COLUMN positionorder TO position_order;
ALTER TABLE results RENAME COLUMN fastestlap TO fastest_lap;
ALTER TABLE results RENAME COLUMN fastestlaptime TO fastest_lap_time;
ALTER TABLE results RENAME COLUMN fastestlapspeed TO fastest_lap_speed;
ALTER TABLE results RENAME COLUMN statusid TO status_id;

ALTER TABLE sprint_results RENAME COLUMN sprintresultid TO id;
ALTER TABLE sprint_results RENAME COLUMN raceid TO race_id;
ALTER TABLE sprint_results RENAME COLUMN driverid TO driver_id;
ALTER TABLE sprint_results RENAME COLUMN constructorid TO constructor_id;
ALTER TABLE sprint_results RENAME COLUMN positiontext TO position_text;
ALTER TABLE sprint_results RENAME COLUMN positionorder TO position_order;
ALTER TABLE sprint_results RENAME COLUMN fastestlap TO fastest_lap;
ALTER TABLE sprint_results RENAME COLUMN fastestlaptime TO fastest_lap_time;
ALTER TABLE sprint_results RENAME COLUMN statusid TO status_id;

ALTER TABLE status RENAME COLUMN statusid TO id;
-- +goose StatementEnd

-- +goose Down
SELECT 'noop';
