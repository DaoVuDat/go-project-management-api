-- name: GetUserProfileById :one
SELECT *
FROM user_profile
WHERE id = $1
LIMIT 1;