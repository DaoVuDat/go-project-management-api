-- name: GetUserProfileById :one
SELECT *
FROM user_profile
WHERE id = $1
LIMIT 1;

-- name: CreateUserProfile :one
INSERT INTO user_profile(id, first_name, last_name, image_url)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUserProfile :one
UPDATE user_profile
SET first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name  = COALESCE(sqlc.narg('last_name'), last_name),
    updated_at = $2
WHERE id = $1
RETURNING *;

-- name: UpdateImageUrlUserProfile :one
UPDATE user_profile
SET image_url = $2
WHERE id = $1
RETURNING *;
