-- name: SaveLapTimes :copyfrom
INSERT INTO lap_times (
    race_id,
    driver_id,
    lap,
    position,
    time,
    milliseconds
)
VALUES ($1, $2, $3, $4, $5, $6);
