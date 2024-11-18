-- name: SavePitStops :copyfrom
INSERT INTO pit_stops (
    race_id,
    driver_id,
    stop,
    lap,
    time,
    duration,
    milliseconds
)
VALUES ($1, $2, $3, $4, $5, $6, $7);
