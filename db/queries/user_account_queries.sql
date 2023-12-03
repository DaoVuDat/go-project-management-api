-- name: GetUserNameAccount :one
SELECT *
FROM user_account
WHERE username = $1
LIMIT 1;

-- name: AddUser :one
INSERT INTO user_account (username, password)
VALUES ($1, $2)
RETURNING *;

-- name: ChangeUserType :one
UPDATE user_account
SET type       = $2,
    updated_at = $3
WHERE user_id = $1
RETURNING *;

-- name: ChangeUserStatus :one
UPDATE user_account
SET status     = $2,
    updated_at = $3
WHERE user_id = $1
RETURNING *;