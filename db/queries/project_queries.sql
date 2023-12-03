-- name: GetProject :one
SELECT *
FROM project
WHERE id = $1
LIMIT 1;

-- name: ListProjects :many
SELECT *
FROM project
ORDER BY created_at;

-- name: GetProjectByUser :one
SELECT *
FROM project
WHERE user_profile = $1
ORDER BY name;