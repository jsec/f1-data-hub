-- name: SaveStatuses :copyfrom
INSERT INTO status (id, status)
VALUES ($1, $2);
