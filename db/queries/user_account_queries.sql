-- name: GetUserNameAccount :one
SELECT *
FROM user_account
WHERE username = $1
LIMIT 1;

-- name: AddUserAccount :one
INSERT INTO user_account (username, password)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUserAccount :one
UPDATE user_account
SET status     = COALESCE(sqlc.narg('status'), status),
    type       = COALESCE(sqlc.narg('type'), type),
    updated_at = $2
WHERE user_id = $1
RETURNING *;
