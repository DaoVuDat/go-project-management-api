-- name: GetUserProfileById :one
SELECT *
FROM user_profile
WHERE id = $1
LIMIT 1;

-- name: UpdateUserProfile :one
UPDATE user_profile
SET first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name  = COALESCE(sqlc.narg('last_name'), last_name),
    image_url  = COALESCE(sqlc.narg('image_url'), image_url),
    updated_at = $2
WHERE id = $1
RETURNING *;

-- name: UpdateImageUrlUserProfile: one
UPDATE user_profile
SET image_url = $2
WHERE id = $1
RETURNING *;
