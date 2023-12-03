-- name: GetUserProfileById :one
SELECT *
FROM user_profile
WHERE id = $1
LIMIT 1;

-- name: UpdateUserProfile :one
UPDATE user_profile
SET first_name = $2,
    last_name  = $3,
    image_url  = $4,
    updated_at = $5
WHERE id = $1
RETURNING *;
