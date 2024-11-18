-- name: SaveCircuits :copyfrom
INSERT INTO circuits (
    id,
    ref,
    name,
    location,
    country,
    lat,
    lng,
    alt,
    url
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
