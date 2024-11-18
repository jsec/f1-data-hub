-- name: SaveDrivers :copyfrom
INSERT INTO drivers (
    id,
    ref,
    number,
    code,
    first_name,
    last_name,
    date_of_birth,
    nationality,
    url
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: SaveDriverStandings :copyfrom
INSERT INTO driver_standings (
    id,
    race_id,
    driver_id,
    points,
    position,
    pos_text,
    wins
)
VALUES ($1, $2, $3, $4, $5, $6, $7);
