-- name: SaveConstructors :copyfrom
INSERT INTO constructors (
    id,
    ref,
    name,
    nationality,
    url
)
VALUES ($1, $2, $3, $4, $5);

-- name: SaveConstructorStandings :copyfrom
INSERT INTO constructor_standings (
    id,
    race_id,
    constructor_id,
    points,
    position,
    pos_text,
    wins
)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: SaveConstructorResults :copyfrom
INSERT INTO constructor_results (
    id,
    race_id,
    constructor_id,
    points,
    status
)
VALUES ($1, $2, $3, $4, $5);
