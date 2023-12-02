-- name: GetProject :one
SELECT * FROM project
WHERE id = $1 LIMIT 1;

-- name: ListProjects :many
SELECT * FROM project
ORDER BY name;