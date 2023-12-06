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
SET type       = COALESCE(sqlc.narg('type'), type),
    status     = COALESCE(sqlc.narg('status'), status),
    password   = COALESCE(sqlc.narg('password'), password),
    updated_at = $2
WHERE user_id = $1
RETURNING *;
