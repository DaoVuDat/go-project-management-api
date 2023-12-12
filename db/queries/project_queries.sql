-- name: GetProject :one
SELECT *
FROM project
WHERE id = $1
LIMIT 1;

-- name: CreateProject :one
INSERT INTO project(user_profile, price, paid)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListProjects :many
SELECT *
FROM project
ORDER BY created_at;

-- name: GetProjectByUser :one
SELECT *
FROM project
WHERE user_profile = $1
ORDER BY name;

-- name: UpdateProjectName :one
UPDATE project
SET name        = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    updated_at  = $2
WHERE id = $1
RETURNING *;


-- name: UpdateProjectTimeWorking :one
UPDATE project
SET start_time = COALESCE(sqlc.narg('start_time'), start_time),
    end_time   = COALESCE(sqlc.narg('end_time'), end_time),
    updated_at = $2
WHERE id = $1
RETURNING *;

-- name: UpdateProjectPaid :one
UPDATE project
SET paid   = COALESCE(sqlc.narg('paid'), paid),
    status = COALESCE(sqlc.narg('status'), status)
WHERE id = $1
RETURNING *;
