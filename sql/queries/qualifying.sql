-- name: SaveQualifyingResults :copyfrom
INSERT INTO qualifying (
    id,
    race_id,
    driver_id,
    constructor_id,
    number,
    position,
    q1,
    q2,
    q3
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
);
