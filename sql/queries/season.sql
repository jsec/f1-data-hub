-- name: GetDriverStandingsByYear :many
SELECT
    d.id AS driver_id,
    d.first_name,
    d.last_name,
    sum(re.points) AS points
FROM results AS re
INNER JOIN drivers AS d ON re.driver_id = d.id
INNER JOIN races AS r ON re.race_id = r.id
WHERE r.year = $1
GROUP BY d.id
ORDER BY points DESC;
