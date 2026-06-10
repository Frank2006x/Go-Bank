-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password_hash
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  password_hash = COALESCE(sqlc.narg(password_hash), password_hash),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  email = COALESCE(sqlc.narg(email), email)
WHERE
  username = sqlc.arg(username)
RETURNING *;