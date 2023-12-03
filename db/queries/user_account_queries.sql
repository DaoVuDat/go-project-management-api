-- name: GetUserNameAccount :one
SELECT *
FROM user_account
WHERE username = $1
LIMIT 1;