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

-- name: UpdateProjectTimeWorking :one
UPDATE project
SET start_time = $2,
    end_time   = $3
WHERE id = $1
RETURNING *;

-- name: UpdateProjectPaid :one
UPDATE project
SET paid = $2
WHERE id = $1
RETURNING *;

-- name: UpdateProjectStatus :one
UPDATE project
SET status = $2
WHERE id = $1
RETURNING *;